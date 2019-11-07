package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

/*
Connect opens a connection to the postgres database
*/
func Connect() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Print(err)
	}

	var (
		// local db credential
		/* DbHost= os.Getenv("db_host")
		DbUser= os.Getenv("db_user")
		DbPassword= os.Getenv("db_pass")
		DbName= os.Getenv("db_name") */
		
		// docker compose db credentials 
		DbHost     = os.Getenv("POSTGRES_HOST")
		DbUser     = os.Getenv("POSTGRES_USER")
		DbPassword = os.Getenv("POSTGRES_PASSWORD")
		DbName     = os.Getenv("POSTGRESS_DB")
	)
	fmt.Println("dbinfo", DbHost, DbUser, DbPassword, DbName)
	// open a connection to the database
	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DbHost, DbUser, DbPassword, DbName)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return nil, err

	}
	// do not forget to close the connection
	// defer db.Close()

	/* err = db.Ping()
	if err != nil {
		return nil, err
	} */

	fmt.Println("Postgres Successfully connected!")

	return db, nil

}
