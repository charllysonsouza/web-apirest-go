package controllers

import (
	"web-apirest-go/database"
	"web-apirest-go/models"
)

func RegisterTransfer(originAccountId uint, targetAccountId uint, transfer_value int) error {

	var transaction models.Transaction

	transaction.OriginAccount = originAccountId
	transaction.TargetAccount = targetAccountId
	transaction.Rate = transferRate
	transaction.Type = "transfer"
	transaction.Amount = transfer_value

	db := database.GetDatabase()
	err := db.Create(&transaction).Error
	return err
}

func RegisterWithdraw(accountNumberId uint, withdraw_value int) error {
	var transaction models.Transaction

	transaction.OriginAccount = accountNumberId
	transaction.Amount = withdraw_value
	transaction.Rate = withDrawRate
	transaction.Type = "withdraw"

	db := database.GetDatabase()
	err := db.Create(&transaction).Error
	return err
}

func RegisterDeposit(accountNumberId uint, rate int, deposit_value int) error {
	var transaction models.Transaction

	transaction.OriginAccount = accountNumberId
	transaction.Amount = deposit_value
	transaction.Rate = rate
	transaction.Type = "deposit"

	db := database.GetDatabase()
	err := db.Create(&transaction).Error
	return err
}
