package users

import (
	"ecormmerce-app/ecormmerce-rest-api/pkg/logging"

	"github.com/go-pg/pg/v9"
)

/*
Repository provides user repository operations
*/
type Repository interface {
	AddUser(*User) bool
	GetAllUsers() ([]User, error)
	FindUserByUsername(string) *User
	FindOrAddUser(*User) (*User, error)
	GetAllUserRoles() ([]UserRole, error)
	UpdateUser(user *User) (*User, error)
	DeleteUser(user *User) error
}

type repository struct {
	db *pg.DB
}

var userRepositoryLogging logging.Logging

/*
NewRepository creates a users repository with the necessary dependencies
*/
func NewRepository(db *pg.DB) Repository {
	userRepositoryLogging = logging.New("user_repository: ")
	return &repository{db}

}

/*
AddUser saves user to the user's table
*/
func (r *repository) AddUser(user *User) bool {
	err := r.db.Insert(user)
	if err != nil {
		userRepositoryLogging.Printlog("AddUser_Error", err.Error())
		return false
	}
	return true

}

/*
Update a user's info
*/
func (r *repository) UpdateUser(user *User) (*User, error) {
	_, err := r.db.Model(user).Column("id", "username", "firstname", "middlename", "lastname", "email_work", "phone_work",
		"email_personal", "phone_personal", "gender", "role", "status", "last_login", "updated_by", "updated_at").WherePK().Update()
	if err != nil {
		userRepositoryLogging.Printlog("UpdateUser_Error", err.Error())
		return &User{}, err
	}
	return user, nil

}

/*
FindOrAddUser finds user or saves user if not found to the user's table
*/
func (r *repository) FindOrAddUser(user *User) (*User, error) {
	//TODO set user role
	// users.Role = 1

	_, err := r.db.Model(user).
		Column("id").
		Where("email_work = ?email_work").
		OnConflict("DO NOTHING"). // OnConflict is optional
		Returning("id").
		SelectOrInsert()
	if err != nil {
		userRepositoryLogging.Printlog("FindORAddUser_Error", err.Error())
		return &User{}, err
	}

	return user, nil

}

/*
DeleteUser saves user to the user's table
*/
func (r *repository) DeleteUser(user *User) error {
	_, err := r.db.Model(user).WherePK().Delete()
	if err != nil {
		userRepositoryLogging.Printlog("DeleteUser_Error", err.Error())
		return err
	}
	return nil

}

/*
GetAllUsers returns all users from the user's table
*/
func (r *repository) GetAllUsers() ([]User, error) {
	users := []User{}
	_, err := r.db.Query(&users, `
		SELECT users.id, users.firstname, users.middlename, users.lastname, users.email_work, users.username,
			users.phone_work, users.email_personal, users.phone_personal, users.gender, user_roles.role_name,
			users.role, users.status, users.last_login, users.updated_by, users.updated_at, users.created_at
		FROM users, user_roles
		WHERE users.role = user_roles.id
	`)
	if err != nil {
		userRepositoryLogging.Printlog("GetAllusers_Error", err.Error())
		return nil, err
	}
	return users, nil
}

/*
GetAllUserRoles returns all user roles from the user_roles table
*/
func (r *repository) GetAllUserRoles() ([]UserRole, error) {
	userRoles := []UserRole{}
	err := r.db.Model(&userRoles).
		Column("id", "role_name", "description").
		Select()
	if err != nil {
		userRepositoryLogging.Printlog("GetAlluserRoles_Error", err.Error())
		return nil, err
	}
	return userRoles, nil
}

/*
FindUserByUsername returns user by username
*/
func (r *repository) FindUserByUsername(username string) *User {
	user := &User{}
	err := r.db.Model(user).
		TableExpr("users").
		ColumnExpr("users.id, users.firstname, users.middlename, users.lastname, users.email_work, users.username").
		ColumnExpr("users.phone_work, users.email_personal, users.phone_personal, users.gender").
		ColumnExpr("users.role, users.status, users.last_login, users.updated_by, users.updated_at, users.created_at").
		ColumnExpr("user_roles.role_name").
		Join("JOIN user_roles ON user_roles.id = users.role").
		Where("users.username = ?", username).Select()
	if err != nil {
		userRepositoryLogging.Printlog("FindUserByUsername_Error", err.Error())
	}
	return user
}
