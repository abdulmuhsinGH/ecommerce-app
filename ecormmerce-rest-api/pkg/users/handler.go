package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-pg/pg/v9"
	"github.com/google/uuid"
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
	newUser := User{}

	err := parseBody(&newUser, request)
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
	user := User{}

	uuid, err := uuid.Parse(mux.Vars(request)["id"])
	if err != nil {
		userHandlerLogging.Printlog("User HandleUpdateUser; Error while converting string to uuid:", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while converting string to uuid", nil))
		return
	}
	user.ID = uuid

	err = parseBody(&user, request)
	if err != nil {
		userHandlerLogging.Printlog("User HandleUpdateUser; Error while decoding request body:", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error while decoding request body", nil))
		return
	}
	fmt.Println("id", user.ID)
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

	uuid, err := uuid.Parse(mux.Vars(request)["id"])
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

func parseBody(user *User, request *http.Request) error {
	err := request.ParseForm()
	if err != nil {
		return err
	}

	user.Firstname = request.Form.Get("firstname")
	user.Lastname = request.Form.Get("lastname")
	user.Username = request.Form.Get("username")
	user.Middlename = request.Form.Get("middlename")
	user.EmailWork = request.Form.Get("email_work")
	user.EmailPersonal = request.Form.Get("email_personal")
	user.PhonePersonal = request.Form.Get("phone_personal")
	user.PhoneWork = request.Form.Get("phone_work")
	user.Gender = request.Form.Get("gender")
	userRole, err := uuid.Parse(request.Form.Get("user_role"))
	if err != nil {
		return err
	}
	user.Role = userRole
	userStatus, err := strconv.ParseBool(request.Form.Get("status"))
	if err != nil {
		return err
	}
	user.Status = userStatus
	if request.Method == http.MethodPut {

		userUpdatedBy, err := uuid.Parse(request.Form.Get("updated_by"))
		if err != nil {
			return err
		}
		user.UpdatedBy = userUpdatedBy
	}
	return nil
}

/*
SetupRoutes sets up routes to respective handlers
*/
func (h *Handlers) SetupRoutes(mux *mux.Router) {
	mux.HandleFunc("/api/users/new", userHandlerLogging.Httplog((auth.ValidateToken(h.handleAddUser, authServer)))).Methods("POST")
	mux.HandleFunc("/api/users", userHandlerLogging.Httplog((auth.ValidateToken(h.handleGetUsers, authServer)))).Methods("GET")
	mux.HandleFunc("/api/users/{id}", userHandlerLogging.Httplog((auth.ValidateToken(h.handleUpdateUser, authServer)))).Methods("PUT")
	mux.HandleFunc("/api/users/{id}", userHandlerLogging.Httplog((auth.ValidateToken(h.handleDeleteUser, authServer)))).Methods("DELETE")
	mux.HandleFunc("/api/users/roles", userHandlerLogging.Httplog((auth.ValidateToken(h.handleGetUserRoles, authServer)))).Methods("GET")
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
