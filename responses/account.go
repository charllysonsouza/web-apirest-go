package responses

type NewAccountResponse struct {
	AccountNumber string  `json:"accountNumber"`
	AgencyNumber  string  `json:"agencyNumber"`
	Balance       float64 `json:"balance"`
}
