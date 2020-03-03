package main

import (
	"ecormmerce-rest-api/pkg/logging"
	server "ecormmerce-rest-api/pkg/server"
	postgres "ecormmerce-rest-api/pkg/storage/postgres"
	users "ecormmerce-rest-api/pkg/users"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

/* var (
	GcukCertFile    = os.Getenv("GCUK_CERT_FILE")
	GcukKeyFile     = os.Getenv("GCUK_KEY_FILE")
	GcukServiceAddr = os.Getenv("GCUK_SERVICE_ADDR")
) */

func main() {
	logging := logging.New("ecommerce_api:")

	db, err := postgres.Connect()
	if err != nil {
		logging.PrintFatal("postgres connection failed:", err)
	}

	u := users.NewHandlers(logging, db)

	router := mux.NewRouter()

	u.SetupRoutes(router)
	srv := server.New(router, ":8080")

	logging.Printlog("server_status","starting")
	err = srv.ListenAndServe()
	if err != nil {
		logging.PrintFatal("server failed to start:", err)
		logging.Printlog("server_status","closed")
	}
}
