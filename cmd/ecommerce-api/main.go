package main

import (
	server "ecormmerce-rest-api/pkg/server"
	postgres "ecormmerce-rest-api/pkg/storage/postgres"
	users "ecormmerce-rest-api/pkg/users"
	"log"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

/* var (
	GcukCertFile    = os.Getenv("GCUK_CERT_FILE")
	GcukKeyFile     = os.Getenv("GCUK_KEY_FILE")
	GcukServiceAddr = os.Getenv("GCUK_SERVICE_ADDR")
) */

func main() {
	logger := log.New(os.Stdout, "ecommerce_api ", log.LstdFlags|log.Lshortfile)

	db, err := postgres.Connect()
	if err != nil {
		logger.Fatalf("postgres connection failed: %v", err)
	}

	u := users.NewHandlers(logger, db)

	router := mux.NewRouter()

	u.SetupRoutes(router)
	//m := http.ServeMux()
	srv := server.New(router, ":8080")

	logger.Println("server starting")
	err = srv.ListenAndServe()
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}
