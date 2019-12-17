package main

import (
	server "ecormmerce-rest-api/pkg/auth_server"

	auth "ecormmerce-rest-api/pkg/auth"
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

	a := auth.NewHandlers(logger, )

	router := mux.NewRouter()

	a.SetupRoutes(router)
	//m := http.ServeMux()
	srv := server.New(router, ":9064")

	logger.Println("server starting")

	err := srv.ListenAndServe()
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}
