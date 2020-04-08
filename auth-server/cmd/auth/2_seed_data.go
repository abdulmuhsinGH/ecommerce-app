package main

import (
	"ecormmerce-rest-api/auth-server/pkg/auth"
	"fmt"
	"os"

	"github.com/go-pg/migrations/v7"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		var err error
		count, err := db.Model(&auth.OauthClient{}).SelectAndCount()
		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		if count == 0 {
			fmt.Println("seeding oauth_client...")
			_, err = db.Exec(`INSERT INTO ouath_client VALUES (` + os.Getenv("ADMIN_CLIENT_ID") + ` ,` +
				os.Getenv("ADMIN_CLIENT_SECRET") + ` ,` + os.Getenv("ADMIN_CLIENT_DOMAIN") + ` ,` + ` {} )`)

		}
		return err

	}, func(db migrations.DB) error {
		fmt.Println("truncating my_table...")
		_, err := db.Exec(`TRUNCATE oauth_client`)
		return err
	})
}
