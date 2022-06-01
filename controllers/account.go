package controllers

import (
	"log"
	"net/http"
	"strconv"
	"web-apirest-go/database"
	"web-apirest-go/models"
	"web-apirest-go/responses"
	"web-apirest-go/utils"

	"github.com/gin-gonic/gin"
)

func NewAccount(c *gin.Context) {
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

	var newAccount models.Account
	newAccount.Number = utils.GenerateAccountNumber()
	newAccount.Agency = utils.GenerateAgencyNumber()
	newAccount.ClientId = client.ID

	err = db.Create(&newAccount).Error
	if err != nil {
		log.Fatal("cannot create account")
		return
	}

	// mapping to response
	var response = responses.NewAccountResponse{
		AccountNumber: newAccount.Number,
		AgencyNumber:  newAccount.Agency,
		Balance:       0.0,
	}

	c.IndentedJSON(http.StatusCreated, response)
}
