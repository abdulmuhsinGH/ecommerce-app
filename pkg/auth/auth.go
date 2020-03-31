package auth

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"gopkg.in/oauth2.v3/generates"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

var (
	manager *manage.Manager
)

/*
New for authentication
*/
func New() *server.Server {

	manager = manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// generate jwt access token
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte(os.Getenv("jwt_secret")), jwt.SigningMethodHS512))

	srv := server.NewServer(server.NewConfig(), manager)

	return srv
}
