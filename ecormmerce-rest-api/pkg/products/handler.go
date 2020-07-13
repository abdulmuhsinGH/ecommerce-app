package products

import (
	"encoding/json"
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"

	"github.com/go-pg/pg/v9"
	"gopkg.in/oauth2.v3/server"

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
	fmt.Println("add new products")
	newProduct := Product{}

	err := json.NewDecoder(request.Body).Decode(&newProduct)
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
	fmt.Println("add new products")
	product := Product{}

	err := json.NewDecoder(request.Body).Decode(&product)
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

	uuid, err := uuid.FromString(mux.Vars(request)["id"])
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
	mux.HandleFunc("/api/products/new", productHandlerLogging.Httplog((validateToken(h.handleAddProduct)))).Methods("POST")
	mux.HandleFunc("/api/products", productHandlerLogging.Httplog((validateToken(h.handleGetProducts)))).Methods("GET")
	mux.HandleFunc("/api/products", productHandlerLogging.Httplog((validateToken(h.handleUpdateProduct)))).Methods("PUT")
	mux.HandleFunc("/api/products/{id}", productHandlerLogging.Httplog((validateToken(h.handleDeleteProduct)))).Methods("DELETE")
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
