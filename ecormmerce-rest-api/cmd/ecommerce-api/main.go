package main

import (
	"ecormmerce-app/ecormmerce-rest-api/pkg/logging"
	server "ecormmerce-app/ecormmerce-rest-api/pkg/server"
	postgres "ecormmerce-app/ecormmerce-rest-api/pkg/storage/postgres"
	users "ecormmerce-app/ecormmerce-rest-api/pkg/users"
	"os"

	"github.com/gorilla/mux"
)

/* var (
	GcukCertFile    = os.Getenv("GCUK_CERT_FILE")
	GcukKeyFile     = os.Getenv("GCUK_KEY_FILE")
	GcukServiceAddr = os.Getenv("GCUK_SERVICE_ADDR")
) */

func main() {
	logging := logging.New("ecommerce_api:")

	db, err := postgres.Connect(os.Getenv("DB_NAME"))
	if err != nil {
		logging.PrintFatal("postgres connection failed:", err)
	}
	defer db.Close()

	u := users.NewHandlers(logging, db)

	router := mux.NewRouter()

	u.SetupRoutes(router)
	srv := server.New(router, ":"+os.Getenv("PORT"))

	logging.Printlog("server_status", "starting")
	err = srv.ListenAndServe()
	if err != nil {
		logging.PrintFatal("server failed to start:", err)
		logging.Printlog("server_status", "closed")
	}
}
