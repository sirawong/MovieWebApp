package main

import (
	"log"
	controllerAuth "server/controllers/auth"
	"server/db"
	"server/routes"
)

var (
	defaultPort = ":5000"
)

func main() {

	// Init routes
	router := routes.SetupRouter()

	// Connect to Database
	db.ConnectDB()
	controllerAuth.InitSession()

	log.Fatal(router.Run(defaultPort))
}
