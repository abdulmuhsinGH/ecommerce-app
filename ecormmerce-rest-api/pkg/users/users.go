package users

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

/*
User defines the properties of a user
*/
type User struct {
	ID            uuid.UUID `json:"id"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	Firstname     string    `json:"firstname"`
	Middlename    string    `json:"middlename"`
	Lastname      string    `json:"lastname"`
	EmailWork     string    `json:"email_work"`
	PhoneWork     string    `json:"phone_work"`
	EmailPersonal string    `json:"email_personal"`
	PhonePersonal string    `json:"phone_personal"`
	Gender        string    `json:"gender"`
	Role          int       `json:"role"`
	Status        bool      `json:"status"`
	LastLogin     time.Time `json:"last_login"`
	UpdatedBy     string    `json:"updated_by"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `pg:",soft_delete"`
}

/*
UserRole defines the properties of roles a user can have
*/
type UserRole struct {
	ID          int    `pg:"type:integer;primary_key;" json:"id"`
	RoleName    string `json:"role_name"`
	Description string `json:"description"`
	Comment     string `json:"comment"`
	UpdatedBy   string `json:"updated_by"`
}
