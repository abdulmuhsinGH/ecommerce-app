package productcategories

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"

	"github.com/go-pg/pg/v9"
	"gopkg.in/oauth2.v3/server"

	"ecormmerce-app/ecormmerce-rest-api/pkg/auth"
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
	newProductCategory := ProductCategory{}

	err := parseBody(&newProductCategory, request)
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while decoding productCategory", nil))
		return
	}

	err = productCategoryService.AddProductCategory(&newProductCategory)
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
	productCategory := ProductCategory{}

	err := parseBody(&productCategory, request)
	if err != nil {
		format.Send(response, http.StatusInternalServerError, format.Message(false, "Error occured while decoding body", nil))
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

func parseBody(productCategory *ProductCategory, request *http.Request) error {
	err := request.ParseForm()
	if err != nil {
		return err
	}
	ID, err := strconv.ParseInt(request.Form.Get("id"), 10, 64)
	if err != nil {
		return err
	}
	productCategory.ID = ID
	productCategory.Name = request.Form.Get("name")
	productCategory.Description = request.Form.Get("description")
	if request.Method == "PUT" {
		productCatUpdatedBy, err := uuid.Parse(request.Form.Get("updated_by"))
		if err != nil {
			return err
		}
		productCategory.UpdatedBy = productCatUpdatedBy
	}
	return nil
}

/*
SetupRoutes sets up routes to respective handlers
*/
func (h *Handlers) SetupRoutes(mux *mux.Router) {
	mux.HandleFunc("/api/product-categories/new", productCategoryHandlerLogging.Httplog((auth.ValidateToken(h.handleAddProductCategory, authServer)))).Methods("POST")
	mux.HandleFunc("/api/product-categories", productCategoryHandlerLogging.Httplog((auth.ValidateToken(h.handleGetProductCategory, authServer)))).Methods("GET")
	mux.HandleFunc("/api/product-categories", productCategoryHandlerLogging.Httplog((auth.ValidateToken(h.handleUpdateProductCategory, authServer)))).Methods("PUT")
	mux.HandleFunc("/api/product-categories/{id}", productCategoryHandlerLogging.Httplog((auth.ValidateToken(h.handleDeleteProductCategory, authServer)))).Methods("DELETE")
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
