module ecormmerce-app

go 1.13

replace ecormmerce-app/auth-server v0.0.0 => ./auth-server

replace ecormmerce-app/ecormmerce-rest-api v0.0.0 => ./ecormmerce-rest-api

require (
	github.com/go-pg/migrations/v8 v8.0.1 // indirect
	github.com/go-pg/pg/v10 v10.1.0 // indirect
	github.com/joho/godotenv v1.3.0 // indirect
	gopkg.in/oauth2.v3 v3.12.0 // indirect
)
