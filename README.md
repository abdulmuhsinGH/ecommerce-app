
# An Ecommerce API with Golang

## Docker set up 
___
1. run `go mod init`; to get all dependencies in the local project
2. run `go mod verify`; to verify dependencies
### Before you move to the next steps please make sure you have docker and docker-compose installed
3. run `docker-compose build`
4. run `docker-compose up`

## Local set up without Docker
___
1. Create database `ecommerce_db-prod`
2. Run sql scripts in the `/sql` directory
3. run `go mod init`; to get all dependencies in the local project
2. run `go mod verify`; to verify dependencies
5. run `go build cmd/ecommerce-api/main.go`
6. Open Terminal in the project directory and run `./main`

