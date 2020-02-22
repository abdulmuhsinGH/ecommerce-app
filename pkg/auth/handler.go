package auth

import (
	"ecormmerce-rest-api/pkg/format"
	"ecormmerce-rest-api/pkg/users"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/go-session/session"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"gopkg.in/oauth2.v3/server"
)

var (
	userRepository users.Repository
	srv            *server.Server
)

/*
Handlers define auth
*/
type Handlers struct {
	logger *log.Logger
}

/*
Resp interface for response structure
*/
type Resp map[string]interface{}

/*
Logger handles logs
*/
func (h *Handlers) handleLog(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}
func (h *Handlers) handleToken(response http.ResponseWriter, request *http.Request) {
	err := srv.HandleTokenRequest(response, request)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handlers) handleUserAuthTest(response http.ResponseWriter, request *http.Request) {
	token, err := srv.ValidationBearerToken(request)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	data := map[string]interface{}{
		"expires_in": int64(token.GetAccessCreateAt().Add(token.GetAccessExpiresIn()).Sub(time.Now()).Seconds()),
		"client_id":  token.GetClientID(),
		"user_id":    token.GetUserID(),
	}
	e := json.NewEncoder(response)
	e.SetIndent("", "  ")
	e.Encode(data)
}

/*
HandleAddUser gets data from http request and sends to
*/
func (h *Handlers) handleAuthorize(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Hello World")
	store, err := session.Start(nil, response, request)
	if err != nil {
		format.Send(response, 500, format.Message(false, "Error while starting session", nil))
		//http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(store.Get("ReturnUri"))
	var form url.Values
	if v, ok := store.Get("ReturnUri"); ok {
		form = v.(url.Values)
	}
	request.Form = form

	store.Delete("ReturnUri")
	store.Save()

	err = srv.HandleAuthorizeRequest(response, request)
	if err != nil {
		fmt.Println("e", err.Error())
		format.Send(response, 500, format.Message(false, "Error handling authorization", nil))
	}
}
/* 
HandleUserAuthorize handles user authorization
*/
func (h *Handlers) HandleUserAuthorize(response http.ResponseWriter, request *http.Request) {
	store, err := session.Start(nil, response, request)
	if err != nil {
		return
	}

	uid, ok := store.Get("LoggedInUserID")
	if !ok {
		if request.Form == nil {
			request.ParseForm()
		}

		store.Set("ReturnUri", request.Form)
		store.Save()
		fmt.Println(store.Get("ReturnUri"))
		response.Header().Set("Location", "/auth/login")
		response.WriteHeader(http.StatusFound)
		return
	}

	userID := uid.(string)
	fmt.Println(userID)
	store.Delete("LoggedInUserID")
	store.Save()
	return
}

func (h *Handlers) handleLogin(w http.ResponseWriter, r *http.Request) {
	outputHTML(w, r, "pkg/auth/static/login.html")
}
func (h *Handlers) handlePostLogin(w http.ResponseWriter, r *http.Request) {
	store, err := session.Start(nil, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == "POST" {
		store.Set("LoggedInUserID", "000000")
		store.Save()

		w.Header().Set("Location", "/auth")
		w.WriteHeader(http.StatusFound)
		return
	}
}

func (h *Handlers) handleAuth(w http.ResponseWriter, r *http.Request) {
	store, err := session.Start(nil, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, ok := store.Get("LoggedInUserID"); !ok {
		w.Header().Set("Location", "/auth/login")
		w.WriteHeader(http.StatusFound)
		return
	}

	outputHTML(w, r, "pkg/auth/static/auth.html")
}

// outputHTML renders static html files
func outputHTML(w http.ResponseWriter, req *http.Request, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()
	fi, _ := file.Stat()
	http.ServeContent(w, req, file.Name(), fi.ModTime(), file)
}

/*
Logger handles logs
*/
func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}

/*
SetupRoutes sets up routes to respective handlers
*/
func (h *Handlers) SetupRoutes(mux *mux.Router) {
	mux.HandleFunc("/auth", h.handleLog(h.handleAuth)).Methods("GET")
	mux.HandleFunc("/auth/login", h.handleLog(h.handleLogin)).Methods("GET")
	mux.HandleFunc("/auth/login", h.handleLog(h.handlePostLogin)).Methods("POST")
	mux.HandleFunc("/auth/authorize", h.handleLog(h.handleAuthorize)).Methods("GET","POST")
	mux.HandleFunc("/auth/token", h.handleLog(h.handleToken)).Methods("POST")
	mux.HandleFunc("/auth/test", h.handleLog(h.handleUserAuthTest)).Methods("GET")
}

/*
NewHandlers initiates auth handler
*/
func NewHandlers(logger *log.Logger, db *gorm.DB, authServer *server.Server) *Handlers {
	srv = authServer
	userRepository = users.NewRepository(db)
	authService = NewAuthService(userRepository)
	return &Handlers{
		logger: logger,
	}
}
