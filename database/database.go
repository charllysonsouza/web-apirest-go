package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	USER   = "root"
	PASS   = "admin"
	HOST   = "localhost"
	PORT   = 5432
	DBNAME = "bank"
)

func StartDB() {
	strConn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", HOST, USER, PASS, DBNAME, PORT)
	database, err := gorm.Open(postgres.Open(strConn))
	if err != nil {
		log.Panic("Error connecting to database")
	}
	DB = database
}

func GetDatabase() *gorm.DB {
	return DB
}
