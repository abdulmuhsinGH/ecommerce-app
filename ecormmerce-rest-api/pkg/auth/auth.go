package auth

import (
	"os"

	"github.com/go-redis/redis"
	oredis "gopkg.in/go-oauth2/redis.v3"
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
