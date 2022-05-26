package routes

import (
	"web-apirest-go/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	router := gin.Default()

	router.GET("/movimentacoes", controllers.GetMovimentacoes)
	router.POST("/cria_conta", controllers.CriaConta)
	router.POST("/cria_pf", controllers.CreatePF)
	router.POST("/cria_pj", controllers.CreatePJ)
	router.GET("/cliente/:id", controllers.GetById)
	router.POST("/deposito", controllers.Deposito)

	router.Run()
}
