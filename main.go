package main

import (
	"log"
	"os"

	"project-dimsum/config"
	"project-dimsum/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Ambil kunci Firebase dari Environment Variable (untuk di Render)
	firebaseKeyEnv := os.Getenv("FIREBASE_KEY")

	// 2. Nyalakan Database (Mengirimkan string env jika ada)
	// Catatan: Jika fungsi ConnectFirebase() kamu belum menerima parameter,
	// kamu bisa memodifikasi fungsi tersebut di file config/firebase.go
	config.ConnectFirebase(firebaseKeyEnv)
	defer config.DB.Close()

	// 3. Siapkan Router
	r := gin.Default()
	routes.SetupRoutes(r)

	// 4. Jalankan Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Render akan otomatis menyediakan port sendiri lewat env "PORT"
	}

	log.Printf("Server Dimsum berjalan di port %s", port)
	r.Run(":" + port)
}
