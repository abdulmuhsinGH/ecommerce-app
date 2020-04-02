package auth

import (
	"ecormmerce-rest-api/auth-server/pkg/cors"
	"ecormmerce-rest-api/auth-server/pkg/logging"
	"ecormmerce-rest-api/auth-server/pkg/users"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-pg/pg/v9"
	"github.com/go-session/session"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/generates"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

var (
	authService Service
	manager     *manage.Manager
)

/*
Server for authentication
*/
func Server(db *pg.DB, logging logging.Logging) {
	//logger := log.New(os.Stdout, "ecommerce_api ", log.LstdFlags|log.Lshortfile)
	router := mux.NewRouter()
	userRepository := users.NewRepository(db)
	authService = NewAuthService(userRepository)

	manager = manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte(os.Getenv("jwt_secret")), jwt.SigningMethodHS512))

	clientStore := NewClientStore(db)
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
		fmt.Println(r.FormValue("grant_type"))
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

	log.Fatal(http.ListenAndServe(":9096", cors.CORS(router)))
}

func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	//setClientStore(r)
	fmt.Println("Hello user authorize")

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
		fmt.Println("form", r.Form.Encode())
		w.Header().Set("Location", "/auth/login")
		w.WriteHeader(http.StatusFound)
		return
	}
	userID = uid.(uuid.UUID).String()
	store.Delete("LoggedInUserID")
	store.Save()
	return
}
