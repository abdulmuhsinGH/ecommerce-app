package brands

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-pg/pg/v9"
	"gopkg.in/oauth2.v3/server"

	"ecormmerce-app/ecormmerce-rest-api/pkg/format"
	"ecormmerce-app/ecormmerce-rest-api/pkg/logging"

	"github.com/gorilla/mux"
)

/*
Handlers define brand
*/
type Handlers struct {
}

var (
	brandRepository     Repository
	brandService        Service
	brandHandlerLogging logging.Logging
	authServer          *server.Server
)

/*
Resp interface for response structure
*/
type Resp map[string]interface{}

/*
HandleAddbrand gets data from http request and sends to
*/
func (h *Handlers) handleAddBrand(response http.ResponseWriter, request *http.Request) {
	fmt.Println("add new Brands")
	newBrand := ProductBrand{}

	err := json.NewDecoder(request.Body).Decode(&newBrand)
	if err != nil {
		brandHandlerLogging.Printlog("ProductBrand HandleAddBrand; Error while decoding request body:", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error while decoding request body", nil))
		return
	}

	err = brandService.AddBrand(&newBrand)
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while saving brand", nil))
		return
	}
	format.Send(response, http.StatusOK, format.Message(true, "ProductBrand saved", nil))

}

/*
HandleUpdateBrand gets data from http request and sends to
*/
func (h *Handlers) handleUpdateBrand(response http.ResponseWriter, request *http.Request) {
	fmt.Println("update ProductBrand")
	brand := ProductBrand{}

	err := json.NewDecoder(request.Body).Decode(&brand)
	if err != nil {
		brandHandlerLogging.Printlog("ProductBrand HandleUpdateBrand; Error while decoding request body:", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error while decoding request body", nil))
		return
	}

	err = brandService.UpdateBrand(&brand)
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while updating brand", nil))
		return
	}
	format.Send(response, http.StatusOK, format.Message(true, "ProductBrand updated", nil))

}

/*
HandleDeleteBrand gets data from http request and sends to
*/
func (h *Handlers) handleDeleteBrand(response http.ResponseWriter, request *http.Request) {
	brand := ProductBrand{}

	idStr, status := mux.Vars(request)["id"]
	if !status {
		brandHandlerLogging.Printlog("ProductBrand HandleUpdateBrand; Error getting brand id:", "Could not get id")
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while converting string to int", nil))
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		brandHandlerLogging.Printlog("ProductBrand HandleUpdateBrand; Error while converting string to int:", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while converting string to int", nil))
		return
	}
	brand.ID = id
	err = brandService.DeleteBrand(&brand)
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while deleting brand", nil))
		return
	}
	format.Send(response, http.StatusOK, format.Message(true, "ProductBrand deleted", nil))

}

/*
HandleGetBrands gets data from http request and sends to
*/
func (h *Handlers) handleGetBrands(response http.ResponseWriter, request *http.Request) {

	brands, err := brandService.GetAllBrands()
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "error getting all brands", nil))
		return
	}

	format.Send(response, http.StatusOK, format.Message(true, "All brands", brands)) // respond(response, message(true, "ProductBrand saved"))

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
	mux.HandleFunc("/api/brands/new", brandHandlerLogging.Httplog((validateToken(h.handleAddBrand)))).Methods("POST")
	mux.HandleFunc("/api/brands", brandHandlerLogging.Httplog((validateToken(h.handleGetBrands)))).Methods("GET")
	mux.HandleFunc("/api/brands", brandHandlerLogging.Httplog((validateToken(h.handleUpdateBrand)))).Methods("PUT")
	mux.HandleFunc("/api/brands/{id}", brandHandlerLogging.Httplog((validateToken(h.handleDeleteBrand)))).Methods("DELETE")
}

/*
NewHandlers initiates brand handler
*/
func NewHandlers(logger logging.Logging, db *pg.DB, authServerArg *server.Server) *Handlers {
	brandRepository = NewRepository(db)
	brandService = NewService(brandRepository)
	brandHandlerLogging = logger
	authServer = authServerArg

	return &Handlers{}
}
