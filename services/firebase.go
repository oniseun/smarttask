package services

import (
	"context"
	"encoding/base64"
	"log"
	"os"
	"strings"

	db "cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func FirestoreClient() (*db.Client, context.Context) {
	// connect to firebase

	ctx := context.Background()
	serviceAccount, _ := os.LookupEnv("FIREBASE_SERVICE_ACCOUNT")

	credJson, _ := base64.StdEncoding.DecodeString(serviceAccount)
	sa := option.WithCredentialsJSON([]byte(credJson))

	if strings.HasSuffix(serviceAccount, ".json") { // use serviceAccount json file if filename specified
		log.Print("switch to serviceAccount.json")
		sa = option.WithCredentialsFile(serviceAccount)
	}

	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Firestore connected successfully")
	}

	return client, ctx

}
