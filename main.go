package main

import (
	"tugasakhir/database"
	"tugasakhir/route"
)

func main() {
	// Connect to the database
	database.ConnectDb()

	// Set up and run the router
	r := router.SetupRouter()
	r.Run(":8080")
}
