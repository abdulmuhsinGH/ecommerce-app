package users

import (
	"encoding/json"
	"net/http"

	"github.com/go-pg/pg/v9"
	"gopkg.in/oauth2.v3/server"

	"ecormmerce-rest-api/pkg/auth"
	"ecormmerce-rest-api/pkg/format"
	"ecormmerce-rest-api/pkg/logging"

	"github.com/gorilla/mux"
)

/*
Handlers define user
*/
type Handlers struct {
}

var (
	userRepository     Repository
	userService        Service
	userHandlerLogging logging.Logging
	authServer         *server.Server
)

/*
Resp interface for response structure
*/
type Resp map[string]interface{}

/*
HandleAddUser gets data from http request and sends to
*/
func (h *Handlers) handleAddUser(response http.ResponseWriter, request *http.Request) {

	newUser := User{}

	err := json.NewDecoder(request.Body).Decode(&newUser)
	if err != nil {
		userHandlerLogging.Printlog("User HandleAddUser; Error while decoding request body:", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error while decoding request body", nil))
		return
	}

	err = userService.AddUser(&newUser)
	if err != nil {
		userHandlerLogging.Printlog("User HandleAddUser; Error while saving user: %v", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while saving user", nil))
		return
	}
	format.Send(response, http.StatusOK, format.Message(true, "User saved", nil))

}

/*
HandleGetUsers gets data from http request and sends to
*/
func (h *Handlers) handleGetUsers(response http.ResponseWriter, request *http.Request) {

	users, err := userService.GetAllUsers()
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "error getting all users", nil))
		return
	}
	format.Send(response, http.StatusOK, format.Message(true, "All users", users)) // respond(response, message(true, "User saved"))

}

func validateToken(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := authServer.ValidationBearerToken(r)
		if err != nil {
			format.Send(w, http.StatusBadRequest, format.Message(false, err.Error(), nil))
			//http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		next(w, r)
	})
}

/*
SetupRoutes sets up routes to respective handlers
*/
func (h *Handlers) SetupRoutes(mux *mux.Router) {
	mux.HandleFunc("/api/users/new", userHandlerLogging.Httplog((validateToken(h.handleAddUser)))).Methods("POST")
	mux.HandleFunc("/api/users", userHandlerLogging.Httplog((validateToken(h.handleGetUsers)))).Methods("GET")
}

/*
NewHandlers initiates user handler
*/
func NewHandlers(logger logging.Logging, db *pg.DB) *Handlers {
	userRepository = NewRepository(db)
	userService = NewService(userRepository)
	userHandlerLogging = logger
	authServer = auth.New(db)

	return &Handlers{}
}
