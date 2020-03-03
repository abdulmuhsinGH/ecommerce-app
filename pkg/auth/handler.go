package auth

import (
	"ecormmerce-rest-api/pkg/format"
	logging "ecormmerce-rest-api/pkg/logging"
	"ecormmerce-rest-api/pkg/users"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/go-session/session"
	"github.com/gorilla/mux"
	"gopkg.in/oauth2.v3/server"
)

var (
	srv         *server.Server
	authLogging logging.Logging
)

/*
Handlers define auth
*/
type Handlers struct {
}

func (h *Handlers) handleToken(response http.ResponseWriter, request *http.Request) {
	err := srv.HandleTokenRequest(response, request)
	if err != nil {
		authLogging.Printlog("Error: %s\n", err.Error())
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handlers) handleUserAuthTest(response http.ResponseWriter, request *http.Request) {
	token, err := srv.ValidationBearerToken(request)
	if err != nil {
		authLogging.Printlog("Error: %s\n", err.Error())
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
	store, err := session.Start(nil, response, request)
	if err != nil {
		authLogging.Printlog("Error: %s\n", err.Error())
		format.Send(response, 500, format.Message(false, "Error while starting session", nil))
		return
	}
	var form url.Values
	if v, ok := store.Get("ReturnUri"); ok {
		form = v.(url.Values)
	}
	request.Form = form

	store.Delete("ReturnUri")
	store.Save()

	err = srv.HandleAuthorizeRequest(response, request)
	if err != nil {
		authLogging.Printlog("Error: %s\n", err.Error())
		format.Send(response, 500, format.Message(false, "Error handling authorization", nil))
	}
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
	user, err := authService.Login(r.FormValue("username"), r.FormValue("password"))
	if err != nil {
		authLogging.Printlog("Error: %v", err.Error())
		format.Send(w, http.StatusUnauthorized, format.Message(false, err.Error(), nil))
		return
	}
	store.Set("LoggedInUserID", user.ID)
	store.Save()

	w.Header().Set("Location", "/auth")
	w.WriteHeader(http.StatusFound)
}
func (h *Handlers) handleSignUp(w http.ResponseWriter, r *http.Request) {
	outputHTML(w, r, "pkg/auth/static/signup.html")
}
func (h *Handlers) handlePostSignUp(response http.ResponseWriter, request *http.Request) {
	newUser := users.User{}
	body, err := ioutil.ReadAll(request.Body)

	err = json.Unmarshal([]byte(body), &newUser) //NewDecoder(request.Body).Decode(&newUser)
	if err != nil {
		authLogging.Printlog("Error while decoding request body: %v", err.Error())
		format.Send(response, 500, format.Message(false, "Error while decoding request body", nil))
		return
	}
	err = authService.SignUp(newUser)
	if err != nil {
		authLogging.Printlog("Error: %v", err.Error())
		format.Send(response, http.StatusUnauthorized, format.Message(false, err.Error(), nil))
		return
	}

	response.Header().Set("Location", "/auth/login")
	format.Send(response, http.StatusCreated, format.Message(true, "User Created", nil))
}

func (h *Handlers) handleAuth(w http.ResponseWriter, r *http.Request) {
	store, err := session.Start(nil, w, r)
	if err != nil {
		authLogging.Printlog("Error: %v", err.Error())
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
		defer authLogging.Printlog("request processed in %s\n", time.Now().Sub(startTime).String())
		next(w, r)
	}
}

/*
SetupRoutes sets up routes to respective handlers
*/
func (h *Handlers) SetupRoutes(mux *mux.Router) {
	mux.HandleFunc("/auth", authLogging.Httplog(h.handleAuth)).Methods("GET")
	mux.HandleFunc("/auth/login", authLogging.Httplog(h.handleLogin)).Methods("GET")
	mux.HandleFunc("/auth/login", authLogging.Httplog(h.handlePostLogin)).Methods("POST")
	mux.HandleFunc("/auth/signup", authLogging.Httplog(h.handleSignUp)).Methods("GET")
	mux.HandleFunc("/auth/signup", authLogging.Httplog(h.handlePostSignUp)).Methods("POST")
	mux.HandleFunc("/auth/authorize", authLogging.Httplog(h.handleAuthorize)).Methods("GET", "POST")
	mux.HandleFunc("/auth/token", authLogging.Httplog(h.handleToken)).Methods("POST")
	mux.HandleFunc("/auth/test", authLogging.Httplog(h.handleUserAuthTest)).Methods("GET")
}

/*
NewHandlers initiates auth handler
*/
func NewHandlers(logging logging.Logging, db *pg.DB, authServer *server.Server, service Service) *Handlers {
	srv = authServer
	authService = service
	authLogging = logging
	return &Handlers{}
}
