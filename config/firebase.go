package config

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// DB adalah variabel global (Huruf kapital) agar bisa dipakai di controller
var DB *firestore.Client

// Ubah fungsi ini agar menerima parameter string credentialStr
func ConnectFirebase(credentialStr string) {
	ctx := context.Background()
	var app *firebase.App
	var err error
	var opt option.ClientOption

	// 1. Cek dulu apakah ada string kredensial dari Environment Variable (Untuk di Render)
	if credentialStr != "" {
		opt = option.WithCredentialsJSON([]byte(credentialStr))
		app, err = firebase.NewApp(ctx, nil, opt)
		log.Println("Firebase diinisialisasi menggunakan Environment Variable (Production)")

		// 2. Jika env kosong, cek apakah file fisik firebase-key.json ada di lokal (Untuk di Laptop)
	} else if _, errFile := os.Stat("firebase-key.json"); errFile == nil {
		opt = option.WithCredentialsFile("firebase-key.json")
		app, err = firebase.NewApp(ctx, nil, opt)
		log.Println("Firebase diinisialisasi menggunakan file lokal firebase-key.json (Development)")

		// 3. Jika dua-duanya tidak ada (Fallback menggunakan kredensial default bawaan GCP jika ada)
	} else {
		app, err = firebase.NewApp(ctx, nil)
		log.Println("Firebase diinisialisasi menggunakan Google Default Credentials")
	}

	if err != nil {
		log.Fatalf("Gagal membuka Firebase: %v", err)
	}

	// Koneksikan ke Firestore
	DB, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Gagal masuk ke Firestore: %v", err)
	}
}
