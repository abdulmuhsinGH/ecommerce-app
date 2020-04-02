package auth

import (
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
		Addr: "127.0.0.1:6379",
		DB:   15,
	}))

	return server.NewServer(server.NewConfig(), manager)

}
