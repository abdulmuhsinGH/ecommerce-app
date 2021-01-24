package main

import (
	"ecormmerce-app/ecormmerce-rest-api/pkg/auth"
	"ecormmerce-app/ecormmerce-rest-api/pkg/brands"
	"ecormmerce-app/ecormmerce-rest-api/pkg/logging"
	productcategories "ecormmerce-app/ecormmerce-rest-api/pkg/product_categories"
	"ecormmerce-app/ecormmerce-rest-api/pkg/products"
	server "ecormmerce-app/ecormmerce-rest-api/pkg/server"
	postgres "ecormmerce-app/ecormmerce-rest-api/pkg/storage/postgres"
	users "ecormmerce-app/ecormmerce-rest-api/pkg/users"
	"ecormmerce-app/ecormmerce-rest-api/pkg/variants"
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

	usersHandler := users.NewHandlers(logging, db)
	usersHandler.SetupRoutes(router)

	productsHandler := products.NewHandlers(logging, db, authServer)
	productsHandler.SetupRoutes(router)

	brandshandler := brands.NewHandlers(logging, db, authServer)
	brandshandler.SetupRoutes(router)

	productCatHandler := productcategories.NewHandlers(logging, db, authServer)
	productCatHandler.SetupRoutes(router)

	variantsHandler := variants.NewHandlers(logging, db, authServer)
	variantsHandler.SetupRoutes(router)

	srv := server.New(router, ":"+os.Getenv("PORT"))

	logging.Printlog("server_status", "starting on"+os.Getenv("PORT"))
	err = srv.ListenAndServe()
	if err != nil {
		logging.PrintFatal("server failed to start:", err)
		logging.Printlog("server_status", "closed")
	}
}
