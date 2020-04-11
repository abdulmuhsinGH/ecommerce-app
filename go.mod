module ecormmerce-app

go 1.13

require (
	ecormmerce-app/auth-server v0.0.0
	ecormmerce-app/ecormmerce-rest-api v0.0.0
)

replace ecormmerce-app/auth-server v0.0.0 => ./auth-server

replace ecormmerce-app/ecormmerce-rest-api v0.0.0 => ./ecormmerce-rest-api
