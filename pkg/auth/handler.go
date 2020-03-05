package auth

import (
	"context"
	"crypto/rand"
	"ecormmerce-rest-api/pkg/format"
	logging "ecormmerce-rest-api/pkg/logging"
	"ecormmerce-rest-api/pkg/users"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/go-session/session"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
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

var googleOauthConfig *oauth2.Config = &oauth2.Config{
		RedirectURL:  "http://127.0.0.1:9096/auth/google/callback",
		ClientID:     os.Getenv("google_client_id"),
		ClientSecret: os.Getenv("google_client_secret"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

const oauthGoogleURLAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func (h *Handlers) handlePostLoginWithGoogle(w http.ResponseWriter, r *http.Request) {

	// Create oauthState cookie
	oauthState := generateStateOauthCookie(w)
	u := googleOauthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}

func (h *Handlers) handleGoogleAuthCallback(w http.ResponseWriter, r *http.Request) {
	// Read oauthState from Cookie
	oauthState, err := r.Cookie("oauthstate")
	if err != nil {
		authLogging.Printlog("google_oauth_error", err.Error())
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}

	if r.FormValue("state") != oauthState.Value {
		authLogging.Printlog("google_oauth_error", "invalid oauth google state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	data, err := getUserDataFromGoogle(r.FormValue("code"))
	if err != nil {
		authLogging.Printlog("google_oauth_error", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// GetOrCreate User in your db.
	// Redirect or response with a token.
	// More code .....
	err = authService.SignUpViaGoogle(data)
	if err != nil {
		authLogging.Printlog("google_oauth_error", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	//fmt.Fprintf(w, "UserInfo: %s\n", data)
}

func getUserDataFromGoogle(code string) (users.User, error) {
	// Use code to get token and get user info from Google.
	var user users.User
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return users.User{}, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(oauthGoogleURLAPI + token.AccessToken)
	if err != nil {
		return users.User{}, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&user) //ioutil.ReadAll(response.Body)
	if err != nil {
		return users.User{}, fmt.Errorf("failed read response: %s", err.Error())
	}
	return user, nil
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
	mux.HandleFunc("/auth/google/login", authLogging.Httplog(h.handlePostLoginWithGoogle))
	mux.HandleFunc("/auth/google/callback", authLogging.Httplog(h.handleGoogleAuthCallback))
}

/*
NewHandlers initiates auth handler
*/
func NewHandlers(logging logging.Logging, db *pg.DB, authServer *server.Server, service Service) *Handlers {
	srv = authServer
	authService = service
	authLogging = logging
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://127.0.0.1:9096/auth/google/callback",
		ClientID:     os.Getenv("google_client_id"),
		ClientSecret: os.Getenv("google_client_secret"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	return &Handlers{}
}
