package auth

import (
	"os"

	"github.com/go-pg/pg/v9"
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
func New(db *pg.DB) *server.Server {
	manager = manage.NewDefaultManager()

	manager.MapTokenStorage(oredis.NewRedisStore(&redis.Options{
		Addr: os.Getenv("REDIS_SERVER_HOST") + ":" + os.Getenv("REDIS_SERVER_PORT"),
		Password: os.Getenv("REDIS_SERVER_PASS"),
		DB:   15,
	}))

	return server.NewServer(server.NewConfig(), manager)

}
