package auth

import (
	"ecormmerce-app/ecormmerce-rest-api/pkg/format"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	oredis "gopkg.in/go-oauth2/redis.v3"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
)

var (
	manager *manage.Manager
)

/*
New for authentication
*/
func New() *server.Server {
	manager = manage.NewDefaultManager()
	if len(os.Getenv("REDIS_SERVER_PASS")) > 0 {
		manager.MapTokenStorage(oredis.NewRedisStore(&redis.Options{
			Addr:     os.Getenv("REDIS_SERVER_HOST") + ":" + os.Getenv("REDIS_SERVER_PORT"),
			Password: os.Getenv("REDIS_SERVER_PASS"),
			DB:       15,
		}))
	} else {
		manager.MapTokenStorage(oredis.NewRedisStore(&redis.Options{
			Addr: os.Getenv("REDIS_SERVER_HOST") + ":" + os.Getenv("REDIS_SERVER_PORT"),
			DB:   15,
		}))
	}

	return server.NewServer(server.NewConfig(), manager)

}

/*
ValidateToken validates acces token of an http request
*/
func ValidateToken(next http.HandlerFunc, authServer *server.Server) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenInfo, err := authServer.ValidationBearerToken(r)
		if err != nil {
			if err == errors.ErrInvalidAccessToken {
				format.Send(w, http.StatusUnauthorized, format.Message(false, err.Error(), nil))
			} else {
				format.Send(w, http.StatusBadRequest, format.Message(false, err.Error(), nil))
			}

			return
		}
		if r.Method == http.MethodPut {
			r.Form.Set("updated_by", tokenInfo.GetUserID())
		}
		next(w, r)
	})
}
