package postgres

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

/*
Connect opens a connection to the postgres database
*/
func Connect() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Print(err)
	}

	var (
		// local db credential
		DbHost     = os.Getenv("db_host")
		DbUser     = os.Getenv("db_user")
		DbPassword = os.Getenv("db_pass")
		DbName     = os.Getenv("db_name")

		// docker compose db credentials
		/* DbHost     = os.Getenv("POSTGRES_HOST")
		DbUser     = os.Getenv("POSTGRES_USER")
		DbPassword = os.Getenv("POSTGRES_PASSWORD")
		DbName     = os.Getenv("POSTGRESS_DB") */
	)
	fmt.Println("dbinfo", DbHost, DbUser, DbPassword, DbName)
	// open a connection to the database
	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DbHost, DbUser, DbPassword, DbName)
	db, err := gorm.Open("postgres", dbInfo)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	fmt.Println("Postgres Successfully connected!")

	return db, nil

}
