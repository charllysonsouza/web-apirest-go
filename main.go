package main

import (
	"web-apirest-go/database"
	"web-apirest-go/routes"
)

func main() {
	database.StartDB()
	routes.HandleRequests()
}
