package main

import (
	"ecormmerce-rest-api/pkg/auth"
	"ecormmerce-rest-api/pkg/logging"
	"ecormmerce-rest-api/pkg/storage/postgres"
)

func main() {
	//go auth.Client()

	logging := logging.New("ecommerce_auth:")

	db, err := postgres.Connect()
	if err != nil {
		logging.PrintFatal("postgres connection failed:", err)
	}
	defer db.Close()
	go auth.Server(db, logging)
	auth.Client()
}
