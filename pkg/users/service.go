package users

import (
	"golang.org/x/crypto/bcrypt"
)

// Service provides user adding operations.
type Service interface {
	AddUser(User) error
	Login(User)
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
	
	err = s.userRepository.AddUser(user) // error handling omitted for simplicity
	return err

}

/*
Login authenticates user
*/
func (s *service) Login(user User) {
	_ = s.userRepository.AddUser(user) // error handling omitted for simplicity
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
