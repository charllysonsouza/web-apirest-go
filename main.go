package main

import (
	"web-apirest-go/database"
	"web-apirest-go/migration"
	"web-apirest-go/routes"
)

func main() {
	database.StartDB()
	migration.AutoMigrations()
	routes.HandleRequests()
}
