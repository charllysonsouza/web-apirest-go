package controllers

import (
	"web-apirest-go/database"
	"web-apirest-go/models"

	"github.com/google/uuid"
)

func RegisterTransfer(originAccountId uint, targetAccountId uint, transfer_value int) error {

	var transaction models.Transaction

	transaction.TransactionId = uuid.New().String()
	transaction.OriginAccount = originAccountId
	transaction.TargetAccount = targetAccountId
	transaction.Fee = transferFee
	transaction.Type = "transfer"
	transaction.Amount = transfer_value

	db := database.GetDatabase()
	err := db.Create(&transaction).Error
	return err
}

func RegisterWithdraw(accountNumberId uint, withdraw_value int) error {
	var transaction models.Transaction

	transaction.TransactionId = uuid.New().String()
	transaction.OriginAccount = accountNumberId
	transaction.Amount = withdraw_value
	transaction.Fee = withDrawFee
	transaction.Type = "withdraw"

	db := database.GetDatabase()
	err := db.Create(&transaction).Error
	return err
}

func RegisterDeposit(accountNumberId uint, fee int, deposit_value int) error {
	var transaction models.Transaction

	transaction.TransactionId = uuid.New().String()
	transaction.OriginAccount = accountNumberId
	transaction.Amount = deposit_value
	transaction.Fee = fee
	transaction.Type = "deposit"

	db := database.GetDatabase()
	err := db.Create(&transaction).Error
	return err
}
