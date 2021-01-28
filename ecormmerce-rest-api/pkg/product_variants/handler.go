package productvariants

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
Handlers define productVariant
*/
type Handlers struct {
}

var (
	productVariantRepository     Repository
	productVariantService        Service
	productVariantHandlerLogging logging.Logging
	authServer                   *server.Server
)

/*
Resp interface for response structure
*/
type Resp map[string]interface{}

/*
HandleAddProduct gets data from http request and sends to
*/
func (h *Handlers) handleAddProduct(response http.ResponseWriter, request *http.Request) {
	newProduct := ProductVariant{}

	err := parseBody(&newProduct, request)
	if err != nil {
		productVariantHandlerLogging.Printlog("Product Variant HandleAddProduct; Error while decoding request body:", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error while decoding request body", nil))
		return
	}

	err = productVariantService.AddProductVariant(&newProduct)
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while saving productVariant", nil))
		return
	}
	format.Send(response, http.StatusOK, format.Message(true, "Product Variant saved", nil))

}

/*
HandleUpdateProduct gets data from http request and sends to
*/
func (h *Handlers) handleUpdateProduct(response http.ResponseWriter, request *http.Request) {
	productVariant := ProductVariant{}

	uuid, err := uuid.Parse(mux.Vars(request)["id"])
	if err != nil {
		productVariantHandlerLogging.Printlog("Product Variant HandleUpdateProduct; Error while converting string to uuid:", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while converting string to uuid", nil))
		return
	}
	productVariant.ID = uuid

	err = parseBody(&productVariant, request)
	if err != nil {
		productVariantHandlerLogging.Printlog("Product Variant HandleUpdateProduct; Error while decoding request body:", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error while decoding request body", nil))
		return
	}

	err = productVariantService.UpdateProductVariant(&productVariant)
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while updating productVariant", nil))
		return
	}
	format.Send(response, http.StatusOK, format.Message(true, "Product Variant updated", nil))

}

/*
HandleDeleteProduct gets data from http request and sends to
*/
func (h *Handlers) handleDeleteProduct(response http.ResponseWriter, request *http.Request) {
	productVariant := ProductVariant{}

	uuid, err := uuid.Parse(mux.Vars(request)["id"])
	if err != nil {
		productVariantHandlerLogging.Printlog("Product Variant HandleUpdateProduct; Error while converting string to uuid:", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while converting string to uuid", nil))
		return
	}
	productVariant.ID = uuid
	err = productVariantService.DeleteProductVariant(&productVariant)
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while deleting productVariant", nil))
		return
	}
	format.Send(response, http.StatusOK, format.Message(true, "Product Variant deleted", nil))

}

/*
HandleGetProducts gets data from http request and sends to
*/
func (h *Handlers) handleGetProducts(response http.ResponseWriter, request *http.Request) {

	products, err := productVariantService.GetAllProductVariants()
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "error getting all products", nil))
		return
	}

	format.Send(response, http.StatusOK, format.Message(true, "All products", products)) // respond(response, message(true, "ProductVariant saved"))

}

func parseBody(productVariant *ProductVariant, request *http.Request) error {
	err := request.ParseForm()
	if err != nil {
		return err
	}
	productID, err := uuid.Parse(request.Form.Get("product_id"))
	if err != nil {
		return err
	}
	productVariant.ProductID = productID

	productVariant.ProductVariantName = request.Form.Get("product_variant_name")
	productVariant.SKU = request.Form.Get("sku")
	return nil
}

/*
SetupRoutes sets up routes to respective handlers
*/
func (h *Handlers) SetupRoutes(mux *mux.Router) {
	mux.HandleFunc("/api/product-variants/new", productVariantHandlerLogging.Httplog((auth.ValidateToken(h.handleAddProduct, authServer)))).Methods("POST")
	mux.HandleFunc("/api/product-variants", productVariantHandlerLogging.Httplog((auth.ValidateToken(h.handleGetProducts, authServer)))).Methods("GET")
	mux.HandleFunc("/api/product-variants/{id}", productVariantHandlerLogging.Httplog((auth.ValidateToken(h.handleUpdateProduct, authServer)))).Methods("PUT")
	mux.HandleFunc("/api/product-variants/{id}", productVariantHandlerLogging.Httplog((auth.ValidateToken(h.handleDeleteProduct, authServer)))).Methods("DELETE")
}

/*
NewHandlers initiates productVariant handler
*/
func NewHandlers(logger logging.Logging, db *pg.DB, authServerArg *server.Server) *Handlers {
	productVariantRepository = NewRepository(db)
	productVariantService = NewService(productVariantRepository)
	productVariantHandlerLogging = logger
	authServer = authServerArg

	return &Handlers{}
}
