package auth

import (
	"ecormmerce-rest-api/pkg/users"
	"errors"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/oauth2.v3/server"
)

// Service provides user adding operations.
type Service interface {
	Login(string, string) (*users.User, error)
	HashPassword(string) (string, error)
	CheckPasswordHash(string, string) bool
	SignUp(users.User) error
	SignUpViaGoogle(users.User) error
}

type service struct {
	userRepository users.Repository
}

/*
NewAuthService creates a auth service with the necessary dependencies
*/
func NewAuthService(r users.Repository) Service {
	return &service{r}
}

/*
ValidateToken checks if user token is valid and authorises user to access route
*/
func (s *service) ValidateToken(f http.HandlerFunc, srv *server.Server) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := srv.ValidationBearerToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		f.ServeHTTP(w, r)
	})
}

/*
Login authenticates users
*/
func (s *service) Login(username string, password string) (*users.User, error) {
	user := s.userRepository.FindUserByUsername(username)
	if (&users.User{}) == user {
		return &users.User{}, errors.New("user does not exist")
	}
	passwordMatched := s.CheckPasswordHash(password, user.Password)
	if !passwordMatched {
		return &users.User{}, errors.New("password does not match")
	}

	return user, nil
}

/*
SignUp creates a new user
*/
func (s *service) SignUp(user users.User) error {
	var err error
	user.Password, err = s.HashPassword(user.Password)
	if err != nil {
		return err
	}

	status := s.userRepository.AddUser(&user)
	if !status {
		return errors.New("not created")
	}
	return nil

}

/*
SignUpViaGoogle creates a new user
*/
func (s *service) SignUpViaGoogle(user users.User) error {
	var err error
	user.Password, err = s.HashPassword(user.Password)
	if err != nil {
		return err
	}

	status := s.userRepository.FindOrAddUser(&user)
	if !status {
		return errors.New("not created")
	}
	return nil

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
