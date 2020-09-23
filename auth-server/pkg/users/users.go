package users

import (
	"time"

	"github.com/google/uuid"
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
	Role          uuid.UUID `json:"role"`
	Status        bool      `json:"status"`
	LastLogin     time.Time `json:"last_login"`
	UpdatedBy     uuid.UUID `json:"updated_by"`
}

/*
UserRole defines the properties of roles a user can have
*/
type UserRole struct {
	ID          uuid.UUID `pg:"type:integer;primary_key;" json:"id"`
	RoleName    string    `json:"role_name"`
	Description string    `json:"description"`
	Comment     string    `json:"comment"`
	UpdatedBy   uuid.UUID `json:"updated_by"`
}
