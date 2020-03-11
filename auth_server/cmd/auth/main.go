package main

import (
	"ecormmerce-rest-api/auth_server/pkg/auth"
	"ecormmerce-rest-api/auth_server/pkg/logging"
	"ecormmerce-rest-api/auth_server/pkg/storage/postgres"

	"github.com/joho/godotenv"
)

func main() {
	//go auth.Client()

	logging := logging.New("ecommerce_auth:")

	err := godotenv.Load()
	if err != nil {
		logging.PrintFatal(".env file not found %v", err)
	}

	db, err := postgres.Connect()
	if err != nil {
		logging.PrintFatal("postgres connection failed:", err)
	}
	defer db.Close()
	go auth.Server(db, logging)
	auth.Client()
}
