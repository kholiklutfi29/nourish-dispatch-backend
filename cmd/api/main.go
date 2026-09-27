package main

import (
	"log"
	"os"

	"github.com/kholiklutfi29/nourish-dispatch/internal/database"
	"github.com/kholiklutfi29/nourish-dispatch/internal/router"
)

func main() {

	// Database
	db := database.ConnectPostgres()
	defer db.Close() // running after main() finish (will not trigger unless the server stop)

	// Router + appication dependencies
	r := router.SetupRouter(db)

	// port
	port := os.Getenv("APP_PORT")

	// Start server
	log.Println("Server running on :" + port)

	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
