package routes

import (
	"web-apirest-go/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	router := gin.Default()

	router.GET("/client/:id", controllers.GetById)
	router.POST("/client", controllers.NewCliente)
	router.POST("/account/:id", controllers.NewAccount)
	router.PUT("/deposit", controllers.Deposit)
	router.GET("/account/:id", controllers.GetAccountById)
	router.PUT("/withdraw", controllers.WithDraw)
	router.POST("/transfer", controllers.Transfer)
	router.GET("/extract", controllers.GetExtract)
	router.Run()
}
