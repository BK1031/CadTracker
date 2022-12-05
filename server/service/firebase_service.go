package service

import (
	"context"
	"encoding/base64"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
	"log"
	"server/config"
)

var FirebaseAdmin *firebase.App

func InitializeFirebase() {
	decoded, err := base64.StdEncoding.DecodeString(config.FirebaseServiceAccountEncoded)
	if err != nil {
		log.Fatalf("Error decoding service account: %v\n", err)
	}
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: "https://cad-tracker-1031.firebaseio.com",
		ProjectID:   "cad-tracker-1031",
	}
	opt := option.WithCredentialsJSON(decoded)
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}
	FirebaseAdmin = app
}
