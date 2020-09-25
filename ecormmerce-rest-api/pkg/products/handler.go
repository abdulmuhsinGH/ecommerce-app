package products

import (
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
Handlers define product
*/
type Handlers struct {
}

var (
	productRepository     Repository
	productService        Service
	productHandlerLogging logging.Logging
	authServer            *server.Server
)

/*
Resp interface for response structure
*/
type Resp map[string]interface{}

/*
HandleAddProduct gets data from http request and sends to
*/
func (h *Handlers) handleAddProduct(response http.ResponseWriter, request *http.Request) {
	newProduct := Product{}

	err := parseBody(&newProduct, request)
	if err != nil {
		productHandlerLogging.Printlog("Product HandleAddProduct; Error while decoding request body:", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error while decoding request body", nil))
		return
	}

	err = productService.AddProduct(&newProduct)
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while saving product", nil))
		return
	}
	format.Send(response, http.StatusOK, format.Message(true, "Product saved", nil))

}

/*
HandleUpdateProduct gets data from http request and sends to
*/
func (h *Handlers) handleUpdateProduct(response http.ResponseWriter, request *http.Request) {
	product := Product{}

	err := parseBody(&product, request)
	if err != nil {
		productHandlerLogging.Printlog("Product HandleUpdateProduct; Error while decoding request body:", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error while decoding request body", nil))
		return
	}

	err = productService.UpdateProduct(&product)
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while updating product", nil))
		return
	}
	format.Send(response, http.StatusOK, format.Message(true, "Product updated", nil))

}

/*
HandleDeleteProduct gets data from http request and sends to
*/
func (h *Handlers) handleDeleteProduct(response http.ResponseWriter, request *http.Request) {
	product := Product{}

	uuid, err := uuid.Parse(mux.Vars(request)["id"])
	if err != nil {
		productHandlerLogging.Printlog("Product HandleUpdateProduct; Error while converting string to uuid:", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while converting string to uuid", nil))
		return
	}
	product.ID = uuid
	err = productService.DeleteProduct(&product)
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while deleting product", nil))
		return
	}
	format.Send(response, http.StatusOK, format.Message(true, "Product deleted", nil))

}

/*
HandleGetProducts gets data from http request and sends to
*/
func (h *Handlers) handleGetProducts(response http.ResponseWriter, request *http.Request) {

	products, err := productService.GetAllProducts()
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "error getting all products", nil))
		return
	}

	format.Send(response, http.StatusOK, format.Message(true, "All products", products)) // respond(response, message(true, "Product saved"))

}

func parseBody(product *Product, request *http.Request) error {
	err := request.ParseForm()
	if err != nil {
		return err
	}
	ID, err := uuid.Parse(request.Form.Get("id"))
	if err != nil {
		return err
	}
	product.ID = ID
	product.Name = request.Form.Get("name")
	product.Description = request.Form.Get("description")
	productCategory, err := strconv.ParseInt(request.Form.Get("category"), 10, 64)
	if err != nil {
		return err
	}
	product.Category = productCategory
	productBrand, err := strconv.ParseInt(request.Form.Get("brand"), 10, 64)
	if err != nil {
		return err
	}
	product.Brand = productBrand
	if request.Method == "PUT" {
		productUpdatedBy, err := uuid.Parse(request.Form.Get("updated_by"))
		if err != nil {
			return err
		}
		product.UpdatedBy = productUpdatedBy
	}
	return nil
}

/*
SetupRoutes sets up routes to respective handlers
*/
func (h *Handlers) SetupRoutes(mux *mux.Router) {
	mux.HandleFunc("/api/products/new", productHandlerLogging.Httplog((auth.ValidateToken(h.handleAddProduct, authServer)))).Methods("POST")
	mux.HandleFunc("/api/products", productHandlerLogging.Httplog((auth.ValidateToken(h.handleGetProducts, authServer)))).Methods("GET")
	mux.HandleFunc("/api/products", productHandlerLogging.Httplog((auth.ValidateToken(h.handleUpdateProduct, authServer)))).Methods("PUT")
	mux.HandleFunc("/api/products/{id}", productHandlerLogging.Httplog((auth.ValidateToken(h.handleDeleteProduct, authServer)))).Methods("DELETE")
}

/*
NewHandlers initiates product handler
*/
func NewHandlers(logger logging.Logging, db *pg.DB, authServerArg *server.Server) *Handlers {
	productRepository = NewRepository(db)
	productService = NewService(productRepository)
	productHandlerLogging = logger
	authServer = authServerArg

	return &Handlers{}
}
