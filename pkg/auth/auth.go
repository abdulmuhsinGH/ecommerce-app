package auth

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-pg/pg/v9"
	"gopkg.in/oauth2.v3/generates"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
)

var (
	manager *manage.Manager
)

/*
New for authentication
*/
func New(db *pg.DB) *server.Server {

	manager = manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	tokenStore := NewTokenStore(db)

	manager.MapTokenStorage(tokenStore)

	// generate jwt access token
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte(os.Getenv("jwt_secret")), jwt.SigningMethodHS512))

	srv := server.NewServer(server.NewConfig(), manager)

	return srv
}
