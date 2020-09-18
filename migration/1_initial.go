package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("creating table my_table...")
		_, err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table my_table...")
		_, err := db.Exec(`DROP EXTENSION IF EXISTS "uuid-ossp";`)
		return err
	})
}
