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
	GetAllUsers() ([]User, error)
	FindUserByUsername(string) *User
	FindOrAddUser(*User) bool
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
		userLogging.Printlog("AddUser_Error", err.Error())
		return false
	}
	return true

}

/*
FindOrAddUser finds user or saves user if not found to the user's table
*/
func (r *repository) FindOrAddUser(user *User) bool {

	isCreated, err := r.db.Model(user).
		Column("id").
		Where("email_work = ?email").
		OnConflict("DO NOTHING"). // OnConflict is optional
		Returning("id").
		SelectOrInsert()
	if err != nil {
		userLogging.Printlog("FindORAddUser_Error", err.Error())
		return false
	}
	return isCreated

}

/*
GetAllUsers returns all users from the user's table
*/
func (r *repository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.Select(&users)
	if err !=nil {
		userLogging.Printlog("GetAllusers_Error", err.Error())
		return nil, err
	}
	return users, nil
}

/*
GetAllUsers returns all users from the user's table
*/
func (r *repository) FindUserByUsername(username string) *User {
	user := new(User)
	err := r.db.Model(user).Where("username = ?", username).Select()
	if err != nil {
		userLogging.Printlog("FindUserByUsername_Error", err.Error())
	}
	return user
}
