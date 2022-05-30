package responses

import "time"

type TransferResponse struct {
	Number       string        `json:"accountNumber"`
	Agency       string        `json:"agencyNumber"`
	Balance      float64       `json:"balance"`
	Transactions []ExtractInfo `json:"transactions"`
}

type ExtractInfo struct {
	Id    uint      `json:"id"`
	Value float64   `json:"value"`
	Rate  float64   `json:"rate"`
	Type  string    `json:"type"`
	Date  time.Time `json:"date"`
}
