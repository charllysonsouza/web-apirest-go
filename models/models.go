package models

type Cliente struct {
	ID        int    `json:"id"`
	Nome      string `json:"nome"`
	Documento string `json:"documento"`
	Tipo      string `json:"tipo"`
	Email     string `json:"email"`
	Telefone  string `json:"telefone"`
}

type Conta struct {
	ID         int     `json:"id"`
	Numero     string  `json:"numero"`
	Agencia    string  `json:"agencia"`
	Saldo      float64 `json:"saldo"`
	Cliente_id Cliente `json:"cliente_id"`
}

type Movimentacao struct {
	ID            int     `json:"id"`
	Conta_origem  Conta   `json:"conta_origem"`
	Conta_destino Conta   `json:"conta_destino"`
	Valor         float64 `json:"valor"`
	Taxa          float64 `json:"taxa"`
	Tipo          string  `json:"tipo"`
}

type Deposito struct {
	Numero_conta string  `json:"numero_conta"`
	Valor        float64 `json:"valor"`
}
