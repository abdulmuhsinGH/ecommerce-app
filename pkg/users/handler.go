package users

import (
	"ecormmerce-rest-api/pkg/format"
	"ecormmerce-rest-api/pkg/logging"
	"encoding/json"
	"net/http"

	"github.com/go-pg/pg/v9"
	"github.com/gorilla/mux"
)

/*
Handlers define user
*/
type Handlers struct {
}

var userRepository Repository
var userService Service
var userLogging logging.Logging

/*
HandleAddUser gets data from http request and sends to
*/
func (h *Handlers) handleAddUser(response http.ResponseWriter, request *http.Request) {

	newUser := User{}

	err := json.NewDecoder(request.Body).Decode(&newUser)
	if err != nil {
		userLogging.Printlog("User HandleAddUser; Error while decoding request body: %v", err.Error())
		format.Send(response, 400, format.Message(false, "Error while decoding request body", nil))
		return
	}

	err = userService.AddUser(&newUser)

	if err != nil {
		userLogging.Printlog("User HandleAddUser; Error while saving user: %v", err.Error())
		format.Send(response, 400, format.Message(false, "Error while saving user", nil))
		return
	}
	format.Send(response, 200, format.Message(false, "User Saved", nil))

}

/*
HandleGetAllUsers gets data from http request and sends to
*/
func (h *Handlers) handleGetAllUsers(response http.ResponseWriter, request *http.Request) {

	newUser := User{}

	err := json.NewDecoder(request.Body).Decode(&newUser)
	if err != nil {
		format.Send(response, 400, format.Message(false, "Error while decoding request body", nil))
		return
	}

	users, err := userService.GetAllUsers()
	if err != nil {
		format.Send(response, 500, format.Message(false, "Error getting all users", nil))
		return
	}

	if len(users) == 0 {
		format.Send(response, 200, format.Message(false, "No users", nil))
		return
	}
	format.Send(response, 200, format.Message(false, "all Users", users))

}

/*
SetupRoutes sets up routes to respective handlers
*/
func (h *Handlers) SetupRoutes(mux *mux.Router) {
	mux.HandleFunc("/api/users/new", userLogging.Httplog(h.handleAddUser)).Methods("POST")
	mux.HandleFunc("/api/users/all", userLogging.Httplog(h.handleGetAllUsers)).Methods("GET")
}

/*
NewHandlers initiates user handler
*/
func NewHandlers(logging logging.Logging, db *pg.DB) *Handlers {
	userRepository = NewRepository(db)
	userService = NewService(userRepository)
	userLogging = logging
	return &Handlers{}
}
