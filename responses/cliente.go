package responses

type NewClienteResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Document  string `json:"document"`
	Type      string `json:"type"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
}
