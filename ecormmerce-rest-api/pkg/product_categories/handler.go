package product_categories

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
Handlers define productCategory
*/
type Handlers struct {
}

var (
	productCategoryRepository     Repository
	productCategoryService        Service
	productCategoryHandlerLogging logging.Logging
	authServer                    *server.Server
)

/*
Resp interface for response structure
*/
type Resp map[string]interface{}

/*
HandleAddProductCategory gets data from http request and sends to
*/
func (h *Handlers) handleAddProductCategory(response http.ResponseWriter, request *http.Request) {
	fmt.Println("add new productCategory")
	newproductCategory := ProductCategory{}

	err := json.NewDecoder(request.Body).Decode(&newproductCategory)
	if err != nil {
		productCategoryHandlerLogging.Printlog("productCategory HandleAddProductCategory; Error while decoding request body:", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error while decoding request body", nil))
		return
	}

	err = productCategoryService.AddProductCategory(&newproductCategory)
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while saving productCategory", nil))
		return
	}
	format.Send(response, http.StatusOK, format.Message(true, "productCategory saved", nil))

}

/*
HandleUpdateproductCategory gets data from http request and sends to
*/
func (h *Handlers) handleUpdateProductCategory(response http.ResponseWriter, request *http.Request) {
	fmt.Println("update productCategory")
	productCategory := ProductCategory{}

	err := json.NewDecoder(request.Body).Decode(&productCategory)
	if err != nil {
		productCategoryHandlerLogging.Printlog("productCategory HandleUpdateProductCategory; Error while decoding request body:", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error while decoding request body", nil))
		return
	}

	err = productCategoryService.UpdateProductCategory(&productCategory)
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while updating productCategory", nil))
		return
	}
	format.Send(response, http.StatusOK, format.Message(true, "productCategory updated", nil))

}

/*
handleDeleteProductCategory gets data from http request and sends to
*/
func (h *Handlers) handleDeleteProductCategory(response http.ResponseWriter, request *http.Request) {
	productCategory := ProductCategory{}

	idStr, status := mux.Vars(request)["id"]
	if !status {
		productCategoryHandlerLogging.Printlog("productCategory HandleUpdateproductCategory; Error getting productCategory id:", "Could not get id")
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while converting string to int", nil))
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		productCategoryHandlerLogging.Printlog("productCategory HandleUpdateproductCategory; Error while converting string to int:", err.Error())
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while converting string to int", nil))
		return
	}
	productCategory.ID = id
	err = productCategoryService.DeleteProductCategory(&productCategory)
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while deleting productCategory", nil))
		return
	}
	format.Send(response, http.StatusOK, format.Message(true, "productCategory deleted", nil))

}

/*
HandleGetproductCategorys gets data from http request and sends to
*/
func (h *Handlers) handleGetProductCategory(response http.ResponseWriter, request *http.Request) {

	productCategorys, err := productCategoryService.GetAllProductCategories()
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "error getting all productCategories", nil))
		return
	}

	format.Send(response, http.StatusOK, format.Message(true, "All productCategorys", productCategorys)) // respond(response, message(true, "productCategory saved"))

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
	mux.HandleFunc("/api/product-categories/new", productCategoryHandlerLogging.Httplog((validateToken(h.handleAddProductCategory)))).Methods("POST")
	mux.HandleFunc("/api/product-categories", productCategoryHandlerLogging.Httplog((validateToken(h.handleGetProductCategory)))).Methods("GET")
	mux.HandleFunc("/api/product-categories", productCategoryHandlerLogging.Httplog((validateToken(h.handleUpdateProductCategory)))).Methods("PUT")
	mux.HandleFunc("/api/product-categories/{id}", productCategoryHandlerLogging.Httplog((validateToken(h.handleDeleteProductCategory)))).Methods("DELETE")
}

/*
NewHandlers initiates productCategory handler
*/
func NewHandlers(logger logging.Logging, db *pg.DB, authServerArg *server.Server) *Handlers {
	productCategoryRepository = NewRepository(db)
	productCategoryService = NewService(productCategoryRepository)
	productCategoryHandlerLogging = logger
	authServer = authServerArg

	return &Handlers{}
}
