package responses

import (
	"time"
)

type TransferResponse struct {
	Number       string        `json:"accountNumber"`
	Agency       string        `json:"agencyNumber"`
	Balance      float64       `json:"balance"`
	Transactions []ExtractInfo `json:"transactions"`
}

type ExtractInfo struct {
	TransactionsAsOrigin      []ExtractLine        `json:"transactionsAsOrigin"`
	TransactionsAsDestination []ExtractLineDestiny `json:"transactionsAsDestination"`
}

type ExtractLine struct {
	TransactionId string    `json:"transactionId"`
	Value         float64   `json:"value"`
	Fee           float64   `json:"fee"`
	Type          string    `json:"type"`
	Date          time.Time `json:"date"`
}

type ExtractLineDestiny struct {
	TransactionId string    `json:"transactionId"`
	Value         float64   `json:"value"`
	Type          string    `json:"type"`
	Date          time.Time `json:"date"`
}
