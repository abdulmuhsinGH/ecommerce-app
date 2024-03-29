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
		DbHost     = os.Getenv("DB_HOST")
		DbUser     = os.Getenv("DB_USER")
		DbPassword = os.Getenv("DB_PASS")
		DbName     = os.Getenv("DB_NAME")
		DbPort     = os.Getenv("DB_PORT")
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
