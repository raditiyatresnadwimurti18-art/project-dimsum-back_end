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

func ConnectFirebase() {
	ctx := context.Background()
	var app *firebase.App
	var err error

	if _, errFile := os.Stat("firebase-key.json"); errFile == nil {
		opt := option.WithCredentialsFile("firebase-key.json")
		app, err = firebase.NewApp(ctx, nil, opt)
	} else {
		app, err = firebase.NewApp(ctx, nil)
	}

	if err != nil {
		log.Fatalf("Gagal membuka Firebase: %v", err)
	}

	DB, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Gagal masuk ke Firestore: %v", err)
	}
}
