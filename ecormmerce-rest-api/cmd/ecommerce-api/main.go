package main

import (
	"ecormmerce-app/ecormmerce-rest-api/pkg/auth"
	"ecormmerce-app/ecormmerce-rest-api/pkg/logging"
	"ecormmerce-app/ecormmerce-rest-api/pkg/products"
	server "ecormmerce-app/ecormmerce-rest-api/pkg/server"
	postgres "ecormmerce-app/ecormmerce-rest-api/pkg/storage/postgres"
	users "ecormmerce-app/ecormmerce-rest-api/pkg/users"
	"os"

	"github.com/go-pg/pg/v9"
	"github.com/gorilla/mux"
)

/* var (
	GcukCertFile    = os.Getenv("GCUK_CERT_FILE")
	GcukKeyFile     = os.Getenv("GCUK_KEY_FILE")
	GcukServiceAddr = os.Getenv("GCUK_SERVICE_ADDR")
) */

func main() {
	logging := logging.New("ecommerce_api:")

	var (
		// local db credential
		DbHost     = os.Getenv("DB_HOST")
		DbUser     = os.Getenv("DB_USER")
		DbPassword = os.Getenv("DB_PASS")
		DbPort     = os.Getenv("DB_PORT")
		DbName     = os.Getenv("DB_NAME")
	)

	dbInfo := pg.Options{
		Addr:     DbHost + ":" + DbPort,
		User:     DbUser,
		Password: DbPassword,
		Database: DbName,
	}

	db, err := postgres.Connect(dbInfo)
	if err != nil {
		logging.PrintFatal("postgres connection failed:", err)
	}
	defer db.Close()

	authServer := auth.New()
	router := mux.NewRouter()

	u := users.NewHandlers(logging, db)
	u.SetupRoutes(router)

	p := products.NewHandlers(logging, db, authServer)
	p.SetupRoutes(router)

	srv := server.New(router, ":"+os.Getenv("PORT"))

	logging.Printlog("server_status", "starting")
	err = srv.ListenAndServe()
	if err != nil {
		logging.PrintFatal("server failed to start:", err)
		logging.Printlog("server_status", "closed")
	}
}
