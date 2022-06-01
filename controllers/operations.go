package controllers

import (
	"log"
	"math"
	"net/http"
	"web-apirest-go/database"
	"web-apirest-go/models"
	"web-apirest-go/requests"
	"web-apirest-go/responses"
	"web-apirest-go/utils"

	"github.com/gin-gonic/gin"
)

const withDrawFee = 400
const transferFee = 100

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
	tax := int(math.RoundToEven(float64(amountInCents) * 0.01))
	valueToDeposit := amountInCents - tax

	account.Deposit(valueToDeposit)

	err = db.Save(&account).Error
	if err != nil {
		log.Fatal("cannot make deposit")
		return
	}

	err = RegisterDeposit(account.ID, tax, amountInCents)
	if err != nil {
		log.Fatal("cannot register deposit")
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
	if (amountInCents + withDrawFee) > account.Balance {
		c.JSON(http.StatusBadRequest, gin.H{"message": "there is not enough money to withdraw"})
		return
	}

	account.Withdraw(amountInCents + withDrawFee)

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
	if (amountInCents + transferFee) > originAccount.Balance {
		c.JSON(http.StatusBadRequest, gin.H{"message": "there is not enough money to withdraw"})
		return
	}

	originAccount.Withdraw(amountInCents + transferFee)
	targetAccount.Deposit(amountInCents)

	db.Save(&originAccount)
	db.Save(&targetAccount)

	err = RegisterTransfer(originAccount.ID, targetAccount.ID, amountInCents)
	if err != nil {
		log.Fatal("cannot possible register transaction")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "transfer succeed!"})
}

func GetExtract(c *gin.Context) {
	var originTransactions []models.Transaction
	var destinationTransactions []models.Transaction
	var request requests.Extract
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

	err = db.Where("origin_account = ?", account.ID).Find(&originTransactions).Error
	if err != nil {
		log.Fatal("cannot find transactions: " + err.Error())
	}

	err = db.Where("target_account = ?", account.ID).Find(&destinationTransactions).Error
	if err != nil {
		log.Fatal("cannot find transactions: " + err.Error())
	}

	var extractInfo responses.ExtractInfo

	response.Number = account.Number
	response.Agency = account.Agency
	response.Balance = float64(account.Balance) / 100.00

	for _, transaction := range originTransactions {
		var extractLine responses.ExtractLine
		extractLine.TransactionId = transaction.TransactionId
		extractLine.Type = transaction.Type
		extractLine.Fee = float64(transaction.Fee) / 100.00
		extractLine.Value = float64(transaction.Amount) / 100.00
		extractLine.Date = transaction.CreatedAt

		extractInfo.TransactionsAsOrigin = append(extractInfo.TransactionsAsOrigin, extractLine)
	}

	for _, transaction := range destinationTransactions {
		var extractLine responses.ExtractLineDestiny
		extractLine.TransactionId = transaction.TransactionId
		extractLine.Type = transaction.Type
		extractLine.Value = float64(transaction.Amount) / 100.00
		extractLine.Date = transaction.CreatedAt

		extractInfo.TransactionsAsDestination = append(extractInfo.TransactionsAsDestination, extractLine)
	}

	response.Transactions = append(response.Transactions, extractInfo)

	c.JSON(http.StatusOK, response)
}
