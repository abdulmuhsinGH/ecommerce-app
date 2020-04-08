package main

import (
	"fmt"

	"github.com/go-pg/migrations/v7"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table my_table...")
		_, err := db.Exec(`CREATE TABLE IF NOT EXISTS gopg_migrations()`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table my_table...")
		_, err := db.Exec(`DROP TABLE IF EXISTS gopg_migrations`)
		return err
	})
}