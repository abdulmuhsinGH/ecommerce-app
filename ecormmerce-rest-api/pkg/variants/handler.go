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

	err := parseVariant(&newVariant, request)
	if err != nil {
		variantHandlerLogging.Printlog("Variant HandleAddVariant; Error while decoding request body:", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error while decoding request body", nil))
		return
	}

	err = variantService.AddVariant(&newVariant)
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while saving variant", nil))
		return
	}
	format.Send(response, http.StatusOK, format.Message(true, "Variant saved", nil))

}

func (h *Handlers) handleAddVariantValue(response http.ResponseWriter, request *http.Request) {
	newVariantValue := VariantValue{}

	err := parseVariantValue(&newVariantValue, request)
	if err != nil {
		variantHandlerLogging.Printlog("Variant HandleAddVariantValue; Error while decoding request body:", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error while decoding request body", nil))
		return
	}

	err = variantService.AddVariantValue(&newVariantValue)
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while saving varaint value", nil))
		return
	}
	format.Send(response, http.StatusOK, format.Message(true, "Variant Value saved", nil))

}

/*
HandleGetVariants gets data from http request and sends to
*/
func (h *Handlers) handleGetVariants(response http.ResponseWriter, request *http.Request) {

	variants, err := variantService.GetAllVariants()
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "error getting all products", nil))
		return
	}

	format.Send(response, http.StatusOK, format.Message(true, "All variants", variants)) // respond(response, message(true, "Variant saved"))

}

/*
HandleGetVariantValues gets data from http request and sends to
*/
func (h *Handlers) handleGetVariantValues(response http.ResponseWriter, request *http.Request) {

	variantValues, err := variantService.GetAllVariantValues()
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "error getting all variant values", nil))
		return
	}

	format.Send(response, http.StatusOK, format.Message(true, "All variant values", variantValues)) // respond(response, message(true, "Variant saved"))

}

func parseVariant(variant *Variant, request *http.Request) error {
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

func parseVariantValue(variantValue *VariantValue, request *http.Request) error {
	err := request.ParseForm()
	if err != nil {
		return err
	}
	variantValue.VariantValueName = request.Form.Get("name")

	return nil
}

/*
SetupRoutes sets up routes to respective handlers
*/
func (h *Handlers) SetupRoutes(mux *mux.Router) {
	mux.HandleFunc("/api/variant/new", variantHandlerLogging.Httplog((auth.ValidateToken(h.handleAddVariant, authServer)))).Methods("POST")
	mux.HandleFunc("/api/variants", variantHandlerLogging.Httplog((auth.ValidateToken(h.handleGetVariants, authServer)))).Methods("GET")
	mux.HandleFunc("/api/variant-value/new", variantHandlerLogging.Httplog((auth.ValidateToken(h.handleAddVariantValue, authServer)))).Methods("POST")
	mux.HandleFunc("/api/variant-values", variantHandlerLogging.Httplog((auth.ValidateToken(h.handleGetVariantValues, authServer)))).Methods("GET")
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
