package users

import (
	"ecormmerce-app/ecormmerce-rest-api/pkg/logging"
	"errors"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/oauth2.v3/server"
)

// Service provides user adding operations.
type Service interface {
	AddUser(*User) error
	GetAllUsers() ([]User, error)
	Login(string, string) (*User, error)
	HashPassword(string) (string, error)
	CheckPasswordHash(string, string) bool
	ValidateToken(http.HandlerFunc, *server.Server) http.HandlerFunc
	GetAllUserRoles() ([]UserRole, error)
	UpdateUser(user *User) error
	DeleteUser(user *User) error
}

type service struct {
	userRepository Repository
}

var userServiceLogging logging.Logging

/*
NewService creates a users service with the necessary dependencies
*/
func NewService(r Repository) Service {
	userServiceLogging = logging.New("user_service:")
	return &service{r}
}

/*
AddUser creates a new user
*/
func (s *service) AddUser(user *User) error {
	var err error
	if len(user.Password) == 0 {
		user.Password = fmt.Sprintf("%s@%s", user.Username, user.EmailWork)
	}
	user.Password, err = s.HashPassword(user.Password)
	if err != nil {
		userServiceLogging.Printlog("Password Hash Error;", err.Error())
		return err
	}

	user, err = s.userRepository.FindOrAddUser(user)
	if err != nil {
		return errors.New("not created")
	}
	return nil

}

/*
UpdateUser creates a new user
*/
func (s *service) UpdateUser(user *User) error {

	user.UpdatedAt = time.Now().UTC()
	_, err := s.userRepository.UpdateUser(user)
	if err != nil {
		return errors.New("not updated")
	}
	return nil

}

/*
DeleteUser creates a new user
*/
func (s *service) DeleteUser(user *User) error {
	err := s.userRepository.DeleteUser(user)
	if err != nil {
		return errors.New("not deleted")
	}
	return nil

}

/*
GetAllUsers gets all users
*/
func (s *service) GetAllUsers() ([]User, error) {
	users, err := s.userRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

/*
GetAllUserRoles gets all user roles
*/
func (s *service) GetAllUserRoles() ([]UserRole, error) {
	userRoles, err := s.userRepository.GetAllUserRoles()
	if err != nil {
		return nil, err
	}

	return userRoles, nil
}

/*
Login authenticates users
*/
func (s *service) Login(username string, password string) (*User, error) {
	user := s.userRepository.FindUserByUsername(username)
	if (&User{}) == user {
		return &User{}, errors.New("user does not exist")
	}
	passwordMatched := s.CheckPasswordHash(password, user.Password)
	if !passwordMatched {
		return &User{}, errors.New("password does not match")
	}

	return user, nil
}

/*
HashPassword encrypts user password
*/
func (s *service) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

/*
CheckPasswordHash checks if password entered matches the hashed password
*/
func (s *service) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

/*
ValidateToken checks if user token is valid and authorises user to access route
*/
func (s *service) ValidateToken(next http.HandlerFunc, srv *server.Server) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := srv.ValidationBearerToken(r)
		if err != nil {
			userServiceLogging.Printlog("validate_token_error", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		next(w, r)
	})
}
