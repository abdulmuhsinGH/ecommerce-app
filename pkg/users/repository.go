package users

import (
	"github.com/jinzhu/gorm"
)

/*
Repository provides user repository operations
*/
type Repository interface {
	AddUser(User) bool
	GetAllUsers() []User
	FindUserByUsername(string) User
}

type repository struct {
	db *gorm.DB
}

/*
NewRepository creates a users repository with the necessary dependencies
*/
func NewRepository(db *gorm.DB) Repository {

	return &repository{db}

}

/*
AddUser saves user to the user's table
*/
func (r *repository) AddUser(user User) bool {
	r.db.NewRecord(user)
	r.db.Create(&user)

	return !r.db.NewRecord(user)
}

/*
GetAllUsers returns all users from the user's table
*/
func (r *repository) GetAllUsers() []User {
	var users []User
	r.db.Find(&users)
	return users
}

/*
GetAllUsers returns all users from the user's table
*/
func (r *repository) FindUserByUsername(username string) User {
	var user User
	r.db.Where("username = ?", username).Find(&user)
	return user
}
