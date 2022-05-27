package routes

import (
	"web-apirest-go/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	router := gin.Default()

	router.GET("/cliente/:id", controllers.GetById)

	router.Run()
}
