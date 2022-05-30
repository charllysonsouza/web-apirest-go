package requests

type DepositRequest struct {
	NumberAccount string
	Amount        float64
}

type WithdrawRequest struct {
	NumberAccount string
	Amount        float64
}

type Transfer struct {
	NumberAccountOrigin string
	NumberAccountTarget string
	Amount              float64
}

type Extract struct {
	AccountNumber string `json:"accountNumber"`
	AgencyNumber  string `json:"agencyNumber"`
}
