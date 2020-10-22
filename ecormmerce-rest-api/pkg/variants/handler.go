package variants

import (
	"net/http"

	"github.com/go-pg/pg/v9"
	"github.com/google/uuid"
	"gopkg.in/oauth2.v3/server"

	"ecormmerce-app/ecormmerce-rest-api/pkg/auth"
	"ecormmerce-app/ecormmerce-rest-api/pkg/format"
	"ecormmerce-app/ecormmerce-rest-api/pkg/logging"

	"github.com/gorilla/mux"
)

/*
Handlers define variant
*/
type Handlers struct {
}

var (
	variantRepository     Repository
	variantService        Service
	variantHandlerLogging logging.Logging
	authServer            *server.Server
)

/*
Resp interface for response structure
*/
type Resp map[string]interface{}

/*
HandleAddVariant gets data from http request and sends to
*/
func (h *Handlers) handleAddVariant(response http.ResponseWriter, request *http.Request) {
	newVariant := Variant{}

	err := parseBody(&newVariant, request)
	if err != nil {
		variantHandlerLogging.Printlog("Variant HandleAddVariant; Error while decoding request body:", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error while decoding request body", nil))
		return
	}

	err = variantService.AddVariant(&newVariant)
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while saving product", nil))
		return
	}
	format.Send(response, http.StatusOK, format.Message(true, "Product saved", nil))

}

/*
HandleUpdateVariant gets data from http request and sends to
*/
/* func (h *Handlers) handleUpdateVariant(response http.ResponseWriter, request *http.Request) {
	variant := Variant{}

	uuid, err := uuid.Parse(mux.Vars(request)["id"])
	if err != nil {
		variantHandlerLogging.Printlog("Product HandleUpdateProduct; Error while converting string to uuid:", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while converting string to uuid", nil))
		return
	}
	variant.ID = uuid

	err = parseBody(&variant, request)
	if err != nil {
		variantHandlerLogging.Printlog("Variant HandleUpdateProduct; Error while decoding request body:", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error while decoding request body", nil))
		return
	}

	err = variantService.UpdateVariant(&variant)
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while updating variant", nil))
		return
	}
	format.Send(response, http.StatusOK, format.Message(true, "Variant updated", nil))

} */

/*
HandleDeleteProduct gets data from http request and sends to
*/
/* func (h *Handlers) handleDeleteVariant(response http.ResponseWriter, request *http.Request) {
	variant := Variant{}

	uuid, err := uuid.Parse(mux.Vars(request)["id"])
	if err != nil {
		variantHandlerLogging.Printlog("Variant HandleUpdateVariant; Error while converting string to uuid:", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while converting string to uuid", nil))
		return
	}
	variant.ID = uuid
	err = variantService.DeleteVariant(&variant)
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while deleting variant", nil))
		return
	}
	format.Send(response, http.StatusOK, format.Message(true, "Variant deleted", nil))

} */

/*
HandleGetProducts gets data from http request and sends to
*/
func (h *Handlers) handleGetVariants(response http.ResponseWriter, request *http.Request) {

	variants, err := variantService.GetAllVariants()
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "error getting all products", nil))
		return
	}

	format.Send(response, http.StatusOK, format.Message(true, "All variants", variants)) // respond(response, message(true, "Variant saved"))

}

func parseBody(variant *Variant, request *http.Request) error {
	err := request.ParseForm()
	if err != nil {
		return err
	}
	variant.VariantName = request.Form.Get("name")
	variant.VariantDesc = request.Form.Get("description")

	if request.Method == "PUT" {
		variantUpdatedBy, err := uuid.Parse(request.Form.Get("updated_by"))
		if err != nil {
			return err
		}
		variant.UpdatedBy = variantUpdatedBy
	}
	return nil
}

/*
SetupRoutes sets up routes to respective handlers
*/
func (h *Handlers) SetupRoutes(mux *mux.Router) {
	mux.HandleFunc("/api/variant/new", variantHandlerLogging.Httplog((auth.ValidateToken(h.handleAddVariant, authServer)))).Methods("POST")
	mux.HandleFunc("/api/variants", variantHandlerLogging.Httplog((auth.ValidateToken(h.handleGetVariants, authServer)))).Methods("GET")
	/* mux.HandleFunc("/api/variants/{id}", productHandlerLogging.Httplog((auth.ValidateToken(h.handleUpdateVariant, authServer)))).Methods("PUT")
	mux.HandleFunc("/api/variants/{id}", productHandlerLogging.Httplog((auth.ValidateToken(h.handleDeleteVariant, authServer)))).Methods("DELETE") */
}

/*
NewHandlers initiates product handler
*/
func NewHandlers(logger logging.Logging, db *pg.DB, authServerArg *server.Server) *Handlers {
	variantRepository = NewRepository(db)
	variantService = NewService(variantRepository)
	variantHandlerLogging = logger
	authServer = authServerArg
	return &Handlers{}
}
