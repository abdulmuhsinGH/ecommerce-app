package users

import (
	"github.com/go-pg/pg/v9"
)

/*
Repository provides user repository operations
*/
type Repository interface {
	AddUser(*User) bool
	GetAllUsers() []User
	FindUserByUsername(string) *User
}

type repository struct {
	db *pg.DB
}

/*
NewRepository creates a users repository with the necessary dependencies
*/
func NewRepository(db *pg.DB) Repository {

	return &repository{db}

}

/*
AddUser saves user to the user's table
*/
func (r *repository) AddUser(user *User) bool {

	//r.db.NewRecord(user)
	err := r.db.Insert(user)
	if err != nil {
		return false
	}
	defer r.db.Close()
	return true
	/* defer r.db.Close()

	return !r.db.NewRecord(user) */
}

/*
GetAllUsers returns all users from the user's table
*/
func (r *repository) GetAllUsers() []User {
	var users []User
	r.db.Select(&users)
	r.db.Close()
	return users
}

/*
GetAllUsers returns all users from the user's table
*/
func (r *repository) FindUserByUsername(username string) *User {
	user := new(User)
	err := r.db.Model(user).Where("username = ?", username).Select()
	if err != nil {
	}
	r.db.Close()
	return user
}
