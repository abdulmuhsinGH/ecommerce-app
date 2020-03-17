package auth

import (
	"ecormmerce-rest-api/auth-server/pkg/logging"
	"ecormmerce-rest-api/auth-server/pkg/users"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-pg/pg/v9"
	"github.com/go-session/session"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/generates"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
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

	// generate jwt access token
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte("00000000"), jwt.SigningMethodHS512))

	clientStore := store.NewClientStore()
	clientStore.Set("222222", &models.Client{
		ID:     "222222",
		Secret: "22222222",
		Domain: "http://127.0.0.1:9094",
	})
	manager.MapClientStorage(clientStore)

	adminClientStore := store.NewClientStore()
	adminClientStore.Set("3333333", &models.Client{
		ID:     "3333333",
		Secret: "22222222",
		Domain: "http://localhost:8080",
	})
	manager.MapClientStorage(adminClientStore)

	srv := server.NewServer(server.NewConfig(), manager)
	srv.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
		user, err := authService.Login(username, password)
		if err != nil {
			return "", err
		}
		return user.ID.String(), nil
	})

	srv.SetUserAuthorizationHandler(userAuthorizeHandler)

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

	log.Fatal(http.ListenAndServe(":9096", router))
}

func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	//store.NewClientStore().GetByID("")
	clientID := r.URL.Query()
	if clientID.Get("client_id") == "3333333" {
		q, err := manager.GetClient("3333333")
		if err != nil {
			adminClientStore := store.NewClientStore()
			adminClientStore.Set("3333333", &models.Client{
				ID:     "3333333",
				Secret: "22222222",
				Domain: "http://localhost:8080",
			})
			manager.MapClientStorage(adminClientStore)
			fmt.Println("authorizrh", err.Error())
		} else {
			fmt.Println("authirze", q.GetDomain())
		}
	} else if clientID.Get("client_id") == "222222" {
		q, err := manager.GetClient("222222")
		if err != nil {
			fmt.Println("authorizrh", err.Error())
			clientStore := store.NewClientStore()
			clientStore.Set("222222", &models.Client{
				ID:     "222222",
				Secret: "22222222",
				Domain: "http://127.0.0.1:9094",
			})
			manager.MapClientStorage(clientStore)
			
		} else {
			fmt.Println("authirze", q.GetDomain())
		}
	}
	

	

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
