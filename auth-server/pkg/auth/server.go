package auth

import (
	"ecormmerce-app/auth-server/pkg/clientstore"
	"ecormmerce-app/auth-server/pkg/cors"
	"ecormmerce-app/auth-server/pkg/logging"
	"ecormmerce-app/auth-server/pkg/users"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/go-redis/redis"
	oredis "gopkg.in/go-oauth2/redis.v3"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-pg/pg/v9"
	"github.com/go-session/session"
	"github.com/gorilla/mux"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/generates"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
)

var (
	authService Service
	manager     *manage.Manager
)

/*
Server for authentication
*/
func Server(db *pg.DB, logging logging.Logging) {

	clientStore := clientstore.New(db)

	//logger := log.New(os.Stdout, "ecommerce_api ", log.LstdFlags|log.Lshortfile)
	router := mux.NewRouter()
	userRepository := users.NewRepository(db)
	authService = NewAuthService(userRepository, clientStore)

	manager = manage.NewDefaultManager()

	tokenConfig := &manage.Config{
		AccessTokenExp:    time.Hour * 24,
		RefreshTokenExp:   time.Hour * 24 * 3,
		IsGenerateRefresh: true,
	}
	manager.SetAuthorizeCodeTokenCfg(tokenConfig)

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

	// create client store for admin dashboard. NB set env variables for ADMIN CLIENT before building
	err := clientStore.Create(clientstore.OauthClient{
		ID:     os.Getenv("ADMIN_CLIENT_ID"),
		Secret: os.Getenv("ADMIN_CLIENT_SECRET"),
		Domain: os.Getenv("ADMIN_CLIENT_DOMAIN"),
		Data:   nil,
	})
	if err != nil {
		logging.Printlog("Error Creating admin client", err.Error())
	}

	manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte(os.Getenv("JWT_SECRET")), jwt.SigningMethodHS512))

	//clientStore.Create
	manager.MapClientStorage(clientStore)

	srv := server.NewServer(server.NewConfig(), manager)
	srv.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
		user, err := authService.Login(username, password)
		if err != nil {
			return "", err
		}
		return user.ID.String(), nil
	})
	srv.SetUserAuthorizationHandler(userAuthorizeHandler)

	srv.SetClientInfoHandler(func(r *http.Request) (clientID, clientSecret string, err error) {
		clientID = r.FormValue("client_id")
		clientSecret = r.FormValue("client_secret")

		if clientID == "" || clientSecret == "" {
			err = errors.ErrAccessDenied
			return
		}

		return
	})

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		logging.Printlog("Internal Error:", err.Error())
		return
	})
	srv.SetResponseErrorHandler(func(re *errors.Response) {
		logging.Printlog("Response Error:", re.Error.Error())
	})

	authHandler := NewHandlers(logging, db, srv, authService)
	authHandler.SetupRoutes(router)

	logging.Printlog("AuthServer", "Server is running at 9096 port.")

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), cors.CORS(router)))
}

func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	store, err := session.Start(nil, w, r)
	if err != nil {
		return
	}

	uid, ok := store.Get("LoggedInUserID")
	if !ok {
		if r.Form == nil {
			r.ParseForm()
		}
		store.Set("ReturnUri", r.Form)
		
		store.Save()

		w.Header().Set("Location", "/auth/login")
		w.WriteHeader(http.StatusFound)
		return
	}
	userID = uid.(uuid.UUID).String()
	store.Delete("LoggedInUserID")
	store.Save()
	return
}
