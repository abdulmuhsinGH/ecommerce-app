package postgres

import (
	"os"

	"github.com/go-pg/pg/v9"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

/*
Connect opens a connection to the postgres database
*/
func Connect() (*pg.DB, error) {
	err := godotenv.Load()
	if err != nil {
	}

	var (
		// local db credential
		DbHost     = os.Getenv("db_host")
		DbUser     = os.Getenv("db_user")
		DbPassword = os.Getenv("db_pass")
		DbName     = os.Getenv("db_name")
		DbPort     = os.Getenv("db_port")

		// docker compose db credentials
		/* DbHost     = os.Getenv("POSTGRES_HOST")
		DbUser     = os.Getenv("POSTGRES_USER")
		DbPassword = os.Getenv("POSTGRES_PASSWORD")
		DbName     = os.Getenv("POSTGRESS_DB") */
	)
	// open a connection to the database
	dbInfo := pg.Options{
		Addr:     DbHost + ":" + DbPort,
		User:     DbUser,
		Password: DbPassword,
		Database: DbName,
	}
	db := pg.Connect(&dbInfo)

	return db, nil

}
