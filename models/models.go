package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Name      string  `json:"name"`
	Document  string  `json:"document"`
	Type      string  `json:"type"`
	Email     string  `json:"email"`
	Telephone string  `json:"telephone"`
	Account   Account `gorm:"ForeignKey:ClientId"`
}

type Account struct {
	gorm.Model
	Number             string        `json:"numero"`
	Agency             string        `json:"agencia"`
	Balance            float64       `json:"saldo"`
	ClientId           uint          `json:"cliente_id"`
	TransactionsOrigin []Transaction `gorm:"ForeignKey:OriginAccount"`
	TransactionsTarget []Transaction `gorm:"ForeignKey:TargetAccount"`
}

type Transaction struct {
	gorm.Model
	OriginAccount uint    `json:"origin_account"`
	TargetAccount uint    `json:"target_account"`
	Amount        float64 `json:"amount"`
	Rate          float64 `json:"rate"`
	Type          string  `json:"type"`
}
