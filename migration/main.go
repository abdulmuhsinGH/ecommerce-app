package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
)

const usageText = `This program runs command on the db. Supported commands are:
  - init - creates version info table in the database
  - up - runs all available migrations.
  - up [target] - runs available migrations up to the target one.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
	- set_version [version] - sets db version without running migrations.
Usage:
	go run *.go <command> [args]
create_sql Usage:
  go run *.go <options> [args]
`

func main() {
	var fileName string
	var isTransaction bool
	flag.BoolVar(&isTransaction, "transaction", true, "set whether the file will be run in a transaction or not")
	flag.StringVar(&fileName, "filename", "new_migration_file", "filename to create sql file. \n Must not be space seperated. use underscore to separate eg. create_new_column")

	flag.Usage = usage
	flag.Parse()
	var (
		// local db credential
		DbHost     = os.Getenv("DB_HOST")
		DbUser     = os.Getenv("DB_USER")
		DbPassword = os.Getenv("DB_PASS")
		DbPort     = os.Getenv("DB_PORT")
		DbName     = os.Getenv("DB_NAME")
	)
	db := pg.Connect(&pg.Options{
		Addr:     DbHost + ":" + DbPort,
		User:     DbUser,
		Password: DbPassword,
		Database: DbName,
	})

	if flag.Arg(0) == "create_sql" {
		fileName = strings.ReplaceAll(fileName, " ", "_")
		err := createSQLFiles(db, fileName, isTransaction)
		if err != nil {
			fmt.Print(usageText)
			exitf(err.Error())
		} else {
			//fmt.Printf("Files Created")
			exitf("Files Created")
		}
	} else {
		//migrations.
		oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
		if err != nil {
			exitf(err.Error())
		}
		if newVersion != oldVersion {
			fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
		} else {
			fmt.Printf("version is %d\n", oldVersion)
		}
	}

}

func usage() {
	fmt.Print(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}

func createSQLFiles(db migrations.DB, filename string, isTransaction bool) error {
	vers, err := migrations.Version(db)
	if err != nil {
		return err
	}
	fileNameUp := ""
	fileNameDown := ""
	if isTransaction {
		fileNameUp = fmt.Sprintf("%d_%s.tx.up.sql", vers, filename)
		fileNameDown = fmt.Sprintf("%d_%s.tx.down.sql", vers, filename)
	} else {
		fileNameUp = fmt.Sprintf("%d_%s.up.sql", vers, filename)
		fileNameDown = fmt.Sprintf("%d_%s.down.sql", vers, filename)
	}

	fileUp, err := os.Create(fileNameUp)
	if err != nil {
		return err
	}
	fileDown, err := os.Create(fileNameDown)
	if err != nil {
		return err
	}
	defer fileUp.Close()
	defer fileDown.Close()

	return nil
}
