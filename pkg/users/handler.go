package users

import (
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
func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}

/*
HandleAddUser gets data from http request and sends to
*/
func (h *Handlers) HandleAddUser(response http.ResponseWriter, request *http.Request) {

	newUser := User{}

	err := json.NewDecoder(request.Body).Decode(&newUser)
	if err != nil {
		h.logger.Printf("User HandleAddUser; Error while decoding request body: %v", err.Error())
		respond(response, message(false, "Error while decoding request body"))
		return
	}

	err = userService.AddUser(newUser)

	if err != nil {
		h.logger.Printf("User HandleAddUser; Error while saving user: %v", err.Error())
		respond(response, message(false, "Error while saving user"))
		return
	}
	respond(response, message(true, "User saved"))

}

/*
HandleGetUsers gets data from http request and sends to
*/
func (h *Handlers) HandleGetUsers(response http.ResponseWriter, request *http.Request) {

	newUser := User{}

	err := json.NewDecoder(request.Body).Decode(&newUser)
	if err != nil {
		respond(response, message(false, "Error while decoding request body"))
		return
	}

	err = userService.AddUser(newUser)

	if err != nil {
		respond(response, message(false, "Error while saving user"))
		return
	}
	respond(response, message(true, "User saved"))

}

func respond(response http.ResponseWriter, data Resp) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(data)
}
func message(status bool, message string) Resp {
	return Resp{"status": status, "message": message}
}

/*
SetupRoutes sets up routes to respective handlers
*/
func (h *Handlers) SetupRoutes(mux *mux.Router) {
	mux.HandleFunc("/api/users/new", h.Logger(h.HandleAddUser)).Methods("POST")
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
