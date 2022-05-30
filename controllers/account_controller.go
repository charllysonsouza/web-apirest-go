package controllers

import (
	"net/http"
	"strconv"
	"web-apirest-go/database"
	"web-apirest-go/models"

	"github.com/gin-gonic/gin"
)

func NewAccount(c *gin.Context) {
	var account models.Account
	var client models.Client
	db := database.GetDatabase()

	id := c.Param("id")
	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID has to be integer"})
		return
	}

	err = db.First(&client, newId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot find client: " + err.Error()})
		return
	}

	err = c.ShouldBindJSON(&account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot bind Json"})
		return
	}

	account.ClientId = uint(newId)

	err = db.Create(&account).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot creat account: account number already exists"})
		return
	}

	c.JSON(http.StatusCreated, account)
}

func GetAccountById(c *gin.Context) {
	var account models.Account

	id := c.Param("id")
	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID has to be integer"})
		return
	}

	db := database.GetDatabase()
	err = db.Find(&account, newId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot find account: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)

}
