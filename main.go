package main

import (
	"github.com/PedroPaiao/rest-api-go/database"
	"github.com/PedroPaiao/rest-api-go/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()

	server.Run()
}
