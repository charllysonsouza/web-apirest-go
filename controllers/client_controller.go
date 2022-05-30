package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"web-apirest-go/database"
	"web-apirest-go/models"
	"web-apirest-go/utils"

	"github.com/gin-gonic/gin"
)

func GetById(c *gin.Context) {
	var client models.Client

	id := c.Param("id")
	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID has to be integer"})
		return
	}

	db := database.GetDatabase()
	err = db.First(&client, newId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot find client: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, client)

}

func NewCliente(c *gin.Context) {
	var client models.Client

	err := c.ShouldBindJSON(&client)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot bind Json"})
		return
	}

	isValid := false
	if strings.EqualFold(client.Type, "PessoaFisica") {
		isValid = utils.ValidateCPF(client.Document)
	} else {
		isValid = utils.ValidateCNPJ(client.Document)
	}

	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "document number is invalid "})
		return
	}

	db := database.GetDatabase()

	err = db.Create(&client).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot creat client: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, client)
}
