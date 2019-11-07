package users

import (
	"database/sql"
)

/* 
Repository provides user repository operations 
*/
type Repository interface {
	AddUser(User) error
}

type repository struct {
	db *sql.DB
}

/* 
NewRepository creates a users repository with the necessary dependencies
 */
 func NewRepository(db *sql.DB) Repository {

	return &repository{db}

}
/*
AddUser saves user to the repository
*/
func (r *repository) AddUser(user User) error {

	const q = `INSERT INTO users(username, password, firstname, lastname, gender) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(q, user.Username, user.Password, user.Firstname, user.Lastname, user.Gender)
	return err
}
