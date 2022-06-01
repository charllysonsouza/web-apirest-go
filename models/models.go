package models

import (
	"time"

	"gorm.io/gorm"
)

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
	Number             string        `gorm:"unique" json:"number"`
	Agency             string        `json:"agency"`
	Balance            int           `json:"balance"`
	ClientId           uint          `json:"client_id"`
	TransactionsOrigin []Transaction `gorm:"ForeignKey:OriginAccount"`
	TransactionsTarget []Transaction `gorm:"ForeignKey:TargetAccount"`
}

type Transaction struct {
	TransactionId string `json:"transactionId"`
	OriginAccount uint   `json:"origin_account"`
	TargetAccount uint   `gorm:"Default:null" json:"target_account"`
	Amount        int    `json:"amount"`
	Fee           int    `json:"fee"`
	Type          string `json:"type"`
	CreatedAt     time.Time
}

func (a *Account) Deposit(amount int) {
	a.Balance += amount
}

func (a *Account) Withdraw(amount int) {
	a.Balance -= amount
}

func (client *Client) FillType() {
	if len(client.Document) == 11 {
		client.Type = "PessoaFisica"
	} else {
		client.Type = "PessoaJuridica"
	}
}
