# An Ecommerce APP with Golang

## Docker set up 
___
1. create `.env` files in the root folder and the three folders using the `.env.template`
2. set all environment variables 
3. open terminal and change directory to project directory ie ecommerce-app
3. run `go mod init`; to get all dependencies in the local project
4. run `go mod verify`; to verify dependencies
### Before you move to the next steps please make sure you have docker and docker-compose installed
5. run `docker-compose -f path/to/dock-compose.development.yml build` for developement environment *OR*
   - run `docker-compose -f path/to/dock-compose.production.yml build` for production environment
6. run `docker-compose -f path/to/dock-compose.development.yml up` for developement environment *OR*
   - run `docker-compose -f path/to/dock-compose.production.yml up` for production environment

## DB Migration
___
### Building the binary file for to run migrations in CLI
1. open terminal and change directory to the project folder
2. Cchange directory to the migration folder
3. Run this command to build the bin file - `go build -o migration_cli main.go`
4. You should add binary file in the folder with the name `migration_cli`

### Creating an sql file for migration
This is for when you want to create a sql file to alter a table, seed data or any change to the the database
1. run this command with the name of the file - `./migrate_cli -filename=<filename> create_sql`. NOTE this must be done while in the migration folder. The filename should not be space separated. eg. `./migrate_cli -filename=add_email_column_to_users_table create_sql`. IT should create two files in the folder in this format - `<migration_version>_filename_tx.down.sql` and `<migration_version>_filename_tx.up.sql`
2. Write the script to add changes to the database in the file that ends with `up.sql` and wrtie the sql script to revert that current change in the file that ends with `down.sql`. Eg. If you were adding a column `created_at` to a table `users`, the file `up.sql` will have the the `ALTER ... ADD COLUMN` statement and the `down.sql` will have the `ALTER ... DROP COLUMN` statement. 
3. Run  `./migrate_cli -h`  for help on  the available commands to run. 


