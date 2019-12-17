package users

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

// Service provides user adding operations.
type Service interface {
	AddUser(User) error
	GetAllUsers() []User
	Login(string, string) (User, error)
	HashPassword(string) (string, error)
	CheckPasswordHash(string, string) bool
}

type service struct {
	userRepository Repository
}

/*
NewService creates a users service with the necessary dependencies
*/
func NewService(r Repository) Service {
	return &service{r}
}

/*
AddUser creates a new user
*/
func (s *service) AddUser(user User) error {
	var err error
	user.Password, err = s.HashPassword(user.Password)
	if err != nil {
		return err
	}

	status := s.userRepository.AddUser(user)
	if !status {
		return errors.New("not created")
	}
	return nil

}

/*
GetAllUsers gets all users
*/
func (s *service) GetAllUsers() []User {
	users := s.userRepository.GetAllUsers()
	return users
}

/*
Login authenticates users
*/
func (s *service) Login(username string, password string) (User, error) {
	user := s.userRepository.FindUserByUsername(username)
	if (User{}) == user {
		return User{}, errors.New("user does not exist")
	}
	passwordMatched := s.CheckPasswordHash(password, user.Password)
	if !passwordMatched {
		return User{}, errors.New("password does not match")
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
