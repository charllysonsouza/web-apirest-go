package migration

import (
	"web-apirest-go/database"
	"web-apirest-go/models"
)

func AutoMigrations() {
	db := database.GetDatabase()
	db.AutoMigrate(&models.Client{})
	db.AutoMigrate(&models.Account{})
	db.AutoMigrate(&models.Transaction{})
}
