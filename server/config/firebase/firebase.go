package firebase

import (
	"context"
	"log"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
	"cloud.google.com/go/storage"
)

// Initialize Firebase Firestore
func InitFirestore() *firestore.Client {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./config/firebase/firebase-service-account.json") // Use your credentials file path
	client, err := firestore.NewClient(ctx, "aarthik-setu", sa)
	if err != nil {
		log.Fatalf("Failed to initialize Firestore: %v", err)
	}
	return client
}

// Initialize Firebase Cloud Storage
func InitStorage() *storage.Client {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./config/firebase/firebase-service-account.json") // Use your credentials file path
	client, err := storage.NewClient(ctx, sa)
	if err != nil {
		log.Fatalf("Failed to initialize Cloud Storage: %v", err)
	}
	return client
}