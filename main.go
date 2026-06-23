package main

import (
	"log"

	"os"

	"project-dimsum/config"

	"project-dimsum/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// 1. Nyalakan Database

	config.ConnectFirebase()

	defer config.DB.Close()

	// 2. Siapkan Router

	r := gin.Default()

	routes.SetupRoutes(r)

	// 3. Jalankan Server

	port := os.Getenv("PORT")

	if port == "" {

		port = "8080"

	}

	log.Printf("Server Dimsum berjalan di port %s", port)

	r.Run(":" + port)

}
