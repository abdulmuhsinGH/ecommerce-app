package products

import (
	"ecormmerce-app/ecormmerce-rest-api/pkg/logging"

	"github.com/go-pg/pg/v9"
	"github.com/google/uuid"
)

/*
Repository provides product repository operations
*/
type Repository interface {
	//AddProduct(*Product) bool
	GetAllProducts() ([]Product, error)
	AddProduct(*Product) (*Product, error)
	UpdateProduct(*Product) (*Product, error)
	DeleteProduct(*Product) error
	GetProductByID(uuid.UUID) (Product, error)
	GetProductsByName(string) ([]Product, error)
}

type repository struct {
	db *pg.DB
}

var productRepositoryLogging logging.Logging

/*
NewRepository creates a product repository with the necessary dependencies
*/
func NewRepository(db *pg.DB) Repository {
	productRepositoryLogging = logging.New("product_repository: ")
	return &repository{db}
}

/*
Update a product's info
ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Category    int       `json:"category"`
	Brand       int       `json:"brand"`
	Cost        float64   `json:"cost"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedBy   int       `json:"updated_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
*/
func (r *repository) UpdateProduct(product *Product) (*Product, error) {
	_, err := r.db.Model(product).Column("id", "name", "category", "brand", "cost", "description", "created_at", "updated_by", "updated_at").WherePK().Update()
	if err != nil {
		productRepositoryLogging.Printlog("UpdateProduct_Error", err.Error())
		return &Product{}, err
	}
	return product, nil

}

/*
FindOrAddProduct finds product or saves product if not found to the product's table
*/
func (r *repository) AddProduct(product *Product) (*Product, error) {

	_, err := r.db.Model(product).
		Returning("id").
		Insert()
	if err != nil {
		productRepositoryLogging.Printlog("FindORAddProduct_Error", err.Error())
		return &Product{}, err
	}

	return product, nil

}

/*
DeleteProduct saves product to the product's table
*/
func (r *repository) DeleteProduct(product *Product) error {
	_, err := r.db.Model(product).WherePK().Delete()
	if err != nil {
		productRepositoryLogging.Printlog("DeleteProduct_Error", err.Error())
		return err
	}
	return nil

}

/*
GetAllProducts returns all products from the product's table
*/
func (r *repository) GetAllProducts() ([]Product, error) {
	var products []Product
	err := r.db.Model(&products).
		Column("id", "name", "category", "brand", "cost", "description", "created_at", "updated_by", "updated_at").
		Select()
	if err != nil {
		productRepositoryLogging.Printlog("GetAllproducts_Error", err.Error())
		return nil, err
	}

	return products, nil
}

/*
GetProductByID returns a product by the id from the product's table
*/
func (r *repository) GetProductByID(ID uuid.UUID) (Product, error) {
	var product Product

	err := r.db.Model(&product).
		Column("id", "name", "category", "brand", "cost", "description", "created_at", "updated_by", "updated_at").
		Where("id = ?", ID).
		Select()

	if err != nil {
		productRepositoryLogging.Printlog("GetAllproducts_Error", err.Error())
		return Product{}, err
	}

	return product, nil
}

/*
GetProductsByName returns a product by the id from the product's table
*/
func (r *repository) GetProductsByName(name string) ([]Product, error) {
	var products []Product
	err := r.db.Model(&products).Where("name like ?", "%"+name+"%").
		Column("id", "name", "category", "brand", "cost", "description", "created_at", "updated_by", "updated_at").
		Select()
	if err != nil {
		productRepositoryLogging.Printlog("GetAllproducts_Error", err.Error())
		return nil, err
	}

	return products, nil
}
