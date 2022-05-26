package controllers

import (
	"net/http"
	"strconv"
	"web-apirest-go/models"
	"web-apirest-go/repositories"

	"github.com/gin-gonic/gin"
)

func CreatePJ(c *gin.Context) {

	var pessoa_pj models.Cliente

	err := c.ShouldBindJSON(&pessoa_pj)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if pessoa_pj.Nome == "" || pessoa_pj.Documento == "" || pessoa_pj.Email == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Mandatory fields missing"})
		return
	}

	repositories.CreatePJ(pessoa_pj)

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Client created"})
}

func CreatePF(c *gin.Context) {

	var pessoa_pj models.Cliente

	err := c.ShouldBindJSON(&pessoa_pj)
	if err != nil {
		return
	}

	if pessoa_pj.Nome == "" || pessoa_pj.Documento == "" || pessoa_pj.Email == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Mandatory fields missing"})
		return
	}

	repositories.CreatePF(pessoa_pj)

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Client created"})
}

func GetMovimentacoes(c *gin.Context) {

}

func CriaConta(c *gin.Context) {

}

func GetById(c *gin.Context) {

	param := c.Param("id")
	newid, _ := strconv.Atoi(param)

	cliente := repositories.GetById(newid)

	c.IndentedJSON(http.StatusOK, cliente)
}

func Deposito(c *gin.Context) {
	var deposito models.Deposito

	err := c.ShouldBindJSON(&deposito)
	if err != nil {
		return
	}

}
