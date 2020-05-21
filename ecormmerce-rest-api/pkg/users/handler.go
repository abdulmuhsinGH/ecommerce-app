package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"

	"github.com/go-pg/pg/v9"
	"gopkg.in/oauth2.v3/server"

	"ecormmerce-app/ecormmerce-rest-api/pkg/auth"
	"ecormmerce-app/ecormmerce-rest-api/pkg/format"
	"ecormmerce-app/ecormmerce-rest-api/pkg/logging"

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
	fmt.Println("add new users")
	newUser := User{}

	err := json.NewDecoder(request.Body).Decode(&newUser)
	if err != nil {
		userHandlerLogging.Printlog("User HandleAddUser; Error while decoding request body:", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error while decoding request body", nil))
		return
	}

	err = userService.AddUser(&newUser)
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while saving user", nil))
		return
	}
	format.Send(response, http.StatusOK, format.Message(true, "User saved", nil))

}

/*
HandleUpdateUser gets data from http request and sends to
*/
func (h *Handlers) handleUpdateUser(response http.ResponseWriter, request *http.Request) {
	fmt.Println("add new users")
	user := User{}

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		userHandlerLogging.Printlog("User HandleUpdateUser; Error while decoding request body:", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error while decoding request body", nil))
		return
	}

	err = userService.UpdateUser(&user)
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while updating user", nil))
		return
	}
	format.Send(response, http.StatusOK, format.Message(true, "User updated", nil))

}

/*
HandleDeleteUser gets data from http request and sends to
*/
func (h *Handlers) handleDeleteUser(response http.ResponseWriter, request *http.Request) {
	user := User{}

	uuid, err := uuid.FromString(mux.Vars(request)["id"])
	if err != nil {
		userHandlerLogging.Printlog("User HandleUpdateUser; Error while converting string to uuid:", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while converting string to uuid", nil))
		return
	}
	user.ID = uuid
	err = userService.DeleteUser(&user)
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while deleting user", nil))
		return
	}
	format.Send(response, http.StatusOK, format.Message(true, "User deleted", nil))

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

/*
HandleGetUserRoles gets data from http request and sends to
*/
func (h *Handlers) handleGetUserRoles(response http.ResponseWriter, request *http.Request) {

	userRoles, err := userService.GetAllUserRoles()
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "error getting all user roles", nil))
		return
	}
	format.Send(response, http.StatusOK, format.Message(true, "All user roles", userRoles)) // respond(response, message(true, "User saved"))

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
	mux.HandleFunc("/api/users", userHandlerLogging.Httplog((validateToken(h.handleUpdateUser)))).Methods("PUT")
	mux.HandleFunc("/api/users/{id}", userHandlerLogging.Httplog((validateToken(h.handleDeleteUser)))).Methods("DELETE")
	mux.HandleFunc("/api/users/roles", userHandlerLogging.Httplog((validateToken(h.handleGetUserRoles)))).Methods("GET")
}

/*
NewHandlers initiates user handler
*/
func NewHandlers(logger logging.Logging, db *pg.DB) *Handlers {
	userRepository = NewRepository(db)
	userService = NewService(userRepository)
	userHandlerLogging = logger
	authServer = auth.New()

	return &Handlers{}
}
