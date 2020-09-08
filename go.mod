module ecormmerce-app

go 1.13

require (
	ecormmerce-app/auth-server v0.0.0 // indirect
	ecormmerce-app/ecormmerce-rest-api v0.0.0 // indirect
	github.com/golang-migrate/migrate v3.5.4+incompatible // indirect
	github.com/golang-migrate/migrate/v4 v4.12.2 // indirect
)

replace ecormmerce-app/auth-server v0.0.0 => ./auth-server

replace ecormmerce-app/ecormmerce-rest-api v0.0.0 => ./ecormmerce-rest-api
