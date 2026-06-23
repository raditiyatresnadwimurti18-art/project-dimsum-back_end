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
	// Tambahkan Middleware CORS ini agar Frontend Web (CORS) diizinkan mengambil data
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	routes.SetupRoutes(r)

	// 3. Jalankan Server

	port := os.Getenv("PORT")

	if port == "" {

		port = "8080"

	}

	log.Printf("Server Dimsum berjalan di port %s", port)

	r.Run(":" + port)

}
