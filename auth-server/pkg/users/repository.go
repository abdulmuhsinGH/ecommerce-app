package users

import (
	"ecormmerce-app/auth-server/pkg/logging"

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
	FindUserRoleByName(string) *UserRole
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
	setDefaultUserRole(r, user)
	userRepositoryLogging.Printlog("user new role", user.Role.String())
	err := r.db.Insert(user)
	if err != nil {
		userRepositoryLogging.Printlog("AddUser_Error", err.Error())
		return false
	}
	return true

}

func setDefaultUserRole(r *repository, user *User) {
	userRepositoryLogging.Printlog("User viewer role id", user.Role.String())
	if len(user.Role.String()) == 0 {
		role := (*repository).FindUserRoleByName(r, "viewer")
		userRepositoryLogging.Printlog("User viewer role", role.RoleName)
		user.Role = role.ID
		userRepositoryLogging.Printlog("user new role", user.Role.String())
	}

}

/*
FindOrAddUser finds user or saves user if not found to the user's table
*/
func (r *repository) FindOrAddUser(user *User) (*User, error) {
	setDefaultUserRole(r, user)
	userRepositoryLogging.Printlog("user new role", user.Role.String())
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
GetAllUsers returns all users from the user's table
*/
func (r *repository) GetAllUsers() ([]User, error) {
	users := []User{}
	err := r.db.Model(&users).Select()
	if err != nil {
		userRepositoryLogging.Printlog("GetAllusers_Repo_Error", err.Error())
		return nil, err
	}
	return users, nil
}

/*
FindUserByUsername returns all users from the user's table
*/
func (r *repository) FindUserByUsername(username string) *User {
	user := new(User)

	err := r.db.Model(user).Where("username = ?", username).Select()
	if err != nil {
		userRepositoryLogging.Printlog("FindUserByUsername_Error", err.Error())
	}
	return user
}

/*
FindUserRoleByName returns all users from the user's table
*/
func (r *repository) FindUserRoleByName(role string) *UserRole {
	userRole := new(UserRole)

	err := r.db.Model(userRole).Where("role_name = ?", role).Select()
	if err != nil {
		userRepositoryLogging.Printlog("FindUserRoleByRoleName_Error", err.Error())
	}
	return userRole
}
