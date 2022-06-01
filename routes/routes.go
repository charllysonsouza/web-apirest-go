package routes

import (
	"web-apirest-go/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	router := gin.Default()

	router.POST("/create-client", controllers.NewCliente)
	router.POST("/create-account/:id", controllers.NewAccount)
	router.POST("/deposit", controllers.Deposit)
	router.POST("/withdraw", controllers.WithDraw)
	router.POST("/transfer", controllers.Transfer)
	router.GET("/extract", controllers.GetExtract)
	router.Run()
}
