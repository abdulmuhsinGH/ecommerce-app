package main

import (
	"ecormmerce-rest-api/pkg/auth"
	"ecormmerce-rest-api/pkg/storage/postgres"
	"log"
	"os"
)

func main() {
	//go auth.Client()

	logger := log.New(os.Stdout, "ecommerce_api ", log.LstdFlags|log.Lshortfile)

	db, err := postgres.Connect()
	if err != nil {
		logger.Fatalf("postgres connection failed: %v", err)
	}
	auth.Server(db)
}
