package users

import (
	"ecormmerce-rest-api/pkg/logging"

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

var userRepositoryLogging logging.Logging

/*
NewRepository creates a users repository with the necessary dependencies
*/
func NewRepository(db *pg.DB) Repository {
	userLogging = logging.New("user_reposiroty: ")
	return &repository{db}

}

/*
AddUser saves user to the user's table
*/
func (r *repository) AddUser(user *User) bool {
	err := r.db.Insert(user)
	if err != nil {
		return false
	}
	return true

}

/*
GetAllUsers returns all users from the user's table
*/
func (r *repository) GetAllUsers() []User {
	var users []User
	r.db.Select(&users)
	return users
}

/*
GetAllUsers returns all users from the user's table
*/
func (r *repository) FindUserByUsername(username string) *User {
	user := new(User)
	err := r.db.Model(user).Where("username = ?", username).Select()
	if err != nil {
		userLogging.Printlog("FindUserByUsername_Error",err.Error())
	}
	return user
}
