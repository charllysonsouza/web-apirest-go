package repositories

import (
	"web-apirest-go/models"
)

func GetById(id int) models.Cliente {
	cliente := models.Cliente{}

	err := row.Scan(&cliente.ID, &cliente.Nome, &cliente.Documento, &cliente.Tipo, &cliente.Email, &cliente.Telefone)
	if err != nil {
		panic(err.Error())
	}

	return cliente
}
