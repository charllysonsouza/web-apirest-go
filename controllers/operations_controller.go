package controllers

import (
	"fmt"
	"log"
	"net/http"
	"web-apirest-go/database"
	"web-apirest-go/models"
	"web-apirest-go/requests"
	"web-apirest-go/responses"
	"web-apirest-go/utils"

	"github.com/gin-gonic/gin"
)

const withDrawRate = 400
const transferRate = 100

func Deposit(c *gin.Context) {
	var deposit requests.DepositRequest
	var account models.Account

	db := database.GetDatabase()

	err := c.ShouldBindJSON(&deposit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data type informed invalid"})
		return
	}

	if deposit.Amount < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value. Value must be positive"})
		return
	}

	err = db.Where("number = ?", deposit.NumberAccount).First(&account).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot find informed account"})
		return
	}

	amountInCents := utils.ConvertToCents(deposit.Amount)
	fmt.Println(amountInCents)
	tax := amountInCents / 100
	fmt.Println(tax)
	valueToDeposit := amountInCents - tax
	fmt.Println(valueToDeposit)

	account.Deposit(valueToDeposit)

	err = db.Save(&account).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "could not make the deposit"})
		return
	}

	err = RegisterDeposit(account.ID, tax, valueToDeposit)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "cannot possible register transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deposit succeed!"})
}

func WithDraw(c *gin.Context) {
	var withdraw requests.WithdrawRequest
	var account models.Account

	db := database.GetDatabase()

	err := c.ShouldBindJSON(&withdraw)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "data type informed invalid"})
		return
	}

	if withdraw.Amount < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value. Value must be positive"})
		return
	}

	err = db.Where("number = ?", withdraw.NumberAccount).First(&account).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot find informed account"})
		return
	}

	amountInCents := utils.ConvertToCents(withdraw.Amount)
	if (amountInCents + withDrawRate) > account.Balance {
		c.JSON(http.StatusBadRequest, gin.H{"message": "there is not enough money to withdraw"})
		return
	}

	account.Withdraw(amountInCents + withDrawRate)

	db.Save(&account)
	RegisterWithdraw(account.ID, amountInCents)

	c.JSON(http.StatusOK, gin.H{"message": "withdraw succeed!"})
}

func Transfer(c *gin.Context) {
	var transfer requests.Transfer
	var originAccount models.Account
	var targetAccount models.Account

	db := database.GetDatabase()

	err := c.ShouldBindJSON(&transfer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot bind Json"})
		return
	}

	if transfer.Amount < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value. Value must be positive"})
		return
	}

	err = db.Where("number = ?", transfer.NumberAccountOrigin).First(&originAccount).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot find informed origin account"})
		return
	}

	err = db.Where("number = ?", transfer.NumberAccountTarget).First(&targetAccount).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot find informed target account"})
		return
	}

	amountInCents := utils.ConvertToCents(transfer.Amount)
	if (amountInCents + transferRate) > originAccount.Balance {
		c.JSON(http.StatusBadRequest, gin.H{"message": "there is not enough money to withdraw"})
		return
	}

	originAccount.Withdraw(amountInCents + transferRate)
	targetAccount.Deposit(amountInCents)

	db.Save(&originAccount)
	db.Save(&targetAccount)

	err = RegisterTransfer(originAccount.ID, targetAccount.ID, amountInCents)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cannot possible register transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "transfer succeed!"})
}

func GetExtract(c *gin.Context) {
	var request requests.Extract
	var transactions []models.Transaction
	var account models.Account
	var response responses.TransferResponse

	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Fatal("cannot find json")
	}

	db := database.GetDatabase()

	err = db.Where("number = ?", request.AccountNumber).First(&account).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot find informed account"})
		return
	}

	err = db.Where("origin_account = ?", account.ID).Find(&transactions).Error
	if err != nil {
		log.Fatal("cannot find transactions: " + err.Error())
	}

	response.Number = account.Number
	response.Agency = account.Agency
	response.Balance = float64(account.Balance)

	for _, transaction := range transactions {
		var extractLine responses.ExtractInfo
		extractLine.Id = transaction.ID
		extractLine.Type = transaction.Type
		extractLine.Rate = float64(transaction.Rate) / 100
		extractLine.Value = float64(transaction.Amount) / 100
		extractLine.Date = transaction.CreatedAt

		response.Transactions = append(response.Transactions, extractLine)
	}

	c.JSON(http.StatusOK, response)
}
