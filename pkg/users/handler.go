package users

import (
	"ecormmerce-rest-api/pkg/format"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

/*
Handlers define user
*/
type Handlers struct {
	logger *log.Logger
}

var userRepository Repository
var userService Service

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

/*
HandleAddUser gets data from http request and sends to
*/
func (h *Handlers) handleAddUser(response http.ResponseWriter, request *http.Request) {

	newUser := User{}

	err := json.NewDecoder(request.Body).Decode(&newUser)
	if err != nil {
		h.logger.Printf("User HandleAddUser; Error while decoding request body: %v", err.Error())
		format.Send(response, 400, format.Message(false, "Error while decoding request body", nil))
		return
	}

	err = userService.AddUser(newUser)

	if err != nil {
		h.logger.Printf("User HandleAddUser; Error while saving user: %v", err.Error())
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
		//respond(response, message(false, "Error while decoding request body"))
		return
	}

	users := userService.GetAllUsers()

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
	mux.HandleFunc("/api/users/new", h.handleLog(h.handleAddUser)).Methods("POST")
	mux.HandleFunc("/api/users/all", h.handleLog(h.handleGetAllUsers)).Methods("GET")
}

/*
NewHandlers initiates user handler
*/
func NewHandlers(logger *log.Logger, db *gorm.DB) *Handlers {
	userRepository = NewRepository(db)
	userService = NewService(userRepository)
	return &Handlers{
		logger: logger,
	}
}
