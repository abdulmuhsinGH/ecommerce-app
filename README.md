# An Ecommerce APP with Golang

## Docker set up 
___
1. create `.env` files in the root folder and the three folders using the `.env.template`
2. set all environment variables 
3. run `go mod init`; to get all dependencies in the local project
4. run `go mod verify`; to verify dependencies
### Before you move to the next steps please make sure you have docker and docker-compose installed
5. run `docker-compose -f path/to/dock-compose.development.yml build` for developement environment *OR*
   - run `docker-compose -f path/to/dock-compose.production.yml build` for production environment
6. run `docker-compose -f path/to/dock-compose.development.yml up` for developement environment *OR*
   - run `docker-compose -f path/to/dock-compose.production.yml up` for production environment

