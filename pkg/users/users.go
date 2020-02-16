package users

import (
	"github.com/jinzhu/gorm"
)

/*
User defines the properties of a user
*/
type User struct {
	gorm.Model
	ID      string `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Middlename string `json:"middlename"`
	Lastname  string `json:"lastname"`
	EmailWork string `json:"email_work"`
	PhoneWork string `json:"phone_work"`
	EmailPersonal string `json:"email_personal"`
	PhonePersonal string `json:"phone_personal"`
	Gender    string `json:"gender"`
	Role int `json:"role"`
	Status bool `json:"status"`
	LastLogin string `json:"last_login"`
	CreatedAt string `json:"created_at"`
	UpdateAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	UpdatedBy string `json:"updated_by"`
}
