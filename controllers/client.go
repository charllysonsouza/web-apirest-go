package controllers

import (
	"log"
	"net/http"
	"web-apirest-go/database"
	"web-apirest-go/models"
	"web-apirest-go/responses"
	"web-apirest-go/utils"

	"github.com/gin-gonic/gin"
)

func NewCliente(c *gin.Context) {
	var client models.Client
	var retrivedClient models.Client

	err := c.ShouldBindJSON(&client)
	if err != nil {
		log.Fatal("cannot bind json")
		return
	}

	if !utils.DocumentIsValid(client.Document) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "document number is invalid "})
		return
	}

	db := database.GetDatabase()

	db.Where("document = ?", client.Document).Find(&retrivedClient)
	if retrivedClient.Document != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "DocumentNumber already registered"})
		return
	}

	client.FillType()

	err = db.Create(&client).Error
	if err != nil {
		log.Fatal("error: cannot create client: " + err.Error())
		return
	}

	// mapping to response
	var newClient = responses.NewClienteResponse{
		ID:        client.ID,
		Name:      client.Name,
		Type:      client.Type,
		Document:  client.Document,
		Email:     client.Email,
		Telephone: client.Telephone,
	}

	c.JSON(http.StatusCreated, newClient)
}
