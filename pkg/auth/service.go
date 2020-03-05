package auth

import (
	"context"
	"crypto/rand"
	"ecormmerce-rest-api/pkg/users"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

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
	GenerateStateOauthCookie(http.ResponseWriter) string
	GetUserDataFromGoogle(string) (users.User, error)
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
GenerateStateOauthCookie genearate a state for verifying request to prevent CSRF
*/
func (s *service) GenerateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)
	/* store, err := session.Start(nil, w, r)
	store.Set("ReturnUri", r.Form)
	store.Save() */

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}
/* 
GetUserDataFromGoogle feches user information from Google api
*/
func (s *service) GetUserDataFromGoogle(code string) (users.User, error) {
	const oauthGoogleURLAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	// Use code to get token and get user info from Google.
	var user users.User
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return users.User{}, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(oauthGoogleURLAPI + token.AccessToken)
	if err != nil {
		return users.User{}, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&user) //ioutil.ReadAll(response.Body)
	if err != nil {
		return users.User{}, fmt.Errorf("failed read response: %s", err.Error())
	}
	return user, nil
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
