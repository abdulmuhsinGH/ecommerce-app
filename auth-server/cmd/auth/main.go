package main

import (
	"ecormmerce-rest-api/auth-server/pkg/auth"
	"ecormmerce-rest-api/auth-server/pkg/logging"
	"ecormmerce-rest-api/auth-server/pkg/storage/postgres"
)

func main() {
	logging := logging.New("ecommerce_auth:")

	db, err := postgres.Connect()
	if err != nil {
		logging.PrintFatal("postgres connection failed:", err)
	}

	defer db.Close()
	auth.Server(db, logging)
}
