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
UpdateProduct Update a product's info
*/
func (r *repository) UpdateProduct(product *Product) (*Product, error) {
	_, err := r.db.Model(product).Column("id", "name", "category", "brand", "description", "updated_by", "updated_at").WherePK().Update()
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
	products := []Product{}
	err := r.db.Model(&products).
		ColumnExpr("product.id, product.name, product.category, product.brand, product.description, product.created_at, product.updated_at").
		ColumnExpr("product_brands.name AS brand_name, product_categories.name as category_name").
		Join("JOIN product_brands ON product_brands.id = product.brand").
		Join("JOIN product_categories ON product_categories.id = product.category").
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
	product := Product{}

	err := r.db.Model(&product).
		ColumnExpr("product.id, product.name, product.category, product.brand, product.description, product.created_at, product.updated_at").
		ColumnExpr("product_brands.name AS brand_name, product_categories.name as category_name").
		Join("JOIN product_brands ON product_brands.id = product.brand").
		Join("JOIN product_categories ON product_categories.id = product.category").
		Where("product.id = ?", ID).
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
	products := []Product{}
	err := r.db.Model(&products).
		ColumnExpr("product.id, product.name, product.category, product.brand, product.description, product.created_at, product.updated_at").
		ColumnExpr("product_brands.name AS brand_name, product_categories.name as category_name").
		Join("JOIN product_brands ON product_brands.id = product.brand").
		Join("JOIN product_categories ON product_categories.id = product.category").
		Where("product.name like ?", "%"+name+"%").
		Select()
	if err != nil {
		productRepositoryLogging.Printlog("GetAllproducts_Error", err.Error())
		return nil, err
	}

	return products, nil
}
