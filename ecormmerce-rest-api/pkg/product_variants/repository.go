package productvariants

import (
	"ecormmerce-app/ecormmerce-rest-api/pkg/logging"

	"github.com/go-pg/pg/v9"
	"github.com/google/uuid"
)

/*
Repository provides product variant repository operations
*/
type Repository interface {
	//AddProduct(*ProductVariant) bool
	GetAllProductVariants() ([]ProductVariant, error)
	AddProductVariant(*ProductVariant) (*ProductVariant, error)
	UpdateProductVariant(*ProductVariant) (*ProductVariant, error)
	DeleteProductVariant(*ProductVariant) error
	GetProductVariantByID(uuid.UUID) (ProductVariant, error)
	//GetProductVariantByVariantID(uuid.UUID) (ProductVariant, error)
	GetProductVariantsByName(string) ([]ProductVariant, error)
}

type repository struct {
	db *pg.DB
}

var productVariantRepositoryLogging logging.Logging

/*
NewRepository creates a product variant repository with the necessary dependencies
*/
func NewRepository(db *pg.DB) Repository {
	productVariantRepositoryLogging = logging.New("product_repository: ")
	return &repository{db}
}

/*
type ProductVariant struct {
	ID                  uuid.UUID `json:"id"`
	ProductID           uuid.UUID `json:"product_id"`
	ProductName         string    `json:"product_name" pg:",discard_unknown_columns"`
	SKU                 string    `json:"sku"`
	ProductVariantValue string    `json:"product_variant_name"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	DeletedAt           time.Time `json:"deleted_at"`
}
*/

/*
UpdateProductVariant updates a productVariant's info
*/
func (r *repository) UpdateProductVariant(productVariant *ProductVariant) (*ProductVariant, error) {
	_, err := r.db.Model(productVariant).Column("id", "product_id", "sku", "product_variant_name", "created_at", "updated_at").WherePK().Update()
	if err != nil {
		productVariantRepositoryLogging.Printlog("UpdateProduct_Error", err.Error())
		return &ProductVariant{}, err
	}
	return productVariant, nil

}

/*
AddProductVariant finds product variant or saves product variant if not found to the product variant's table
*/
func (r *repository) AddProductVariant(productVariant *ProductVariant) (*ProductVariant, error) {

	_, err := r.db.Model(productVariant).
		Returning("id").
		Insert()
	if err != nil {
		productVariantRepositoryLogging.Printlog("FindORAddProduct_Error", err.Error())
		return &ProductVariant{}, err
	}

	return productVariant, nil

}

/*
DeleteProductVariant saves product variant to the product variant's table
*/
func (r *repository) DeleteProductVariant(productVariant *ProductVariant) error {
	_, err := r.db.Model(productVariant).WherePK().Delete()
	if err != nil {
		productVariantRepositoryLogging.Printlog("DeleteProduct_Error", err.Error())
		return err
	}
	return nil

}

/*
GetAllProductVariants returns all product variants from the product variant's table
*/
func (r *repository) GetAllProductVariants() ([]ProductVariant, error) {
	productVariants := []ProductVariant{}
	/*
		 ColumnExpr("product_variant.id, product_variant.product_id, product_variant.sku, product_variant.product_variant_name, product_variant.created_at, product_variant.updated_at").
			ColumnExpr("products.name AS product_name").
			Join("JOIN products ON products.id = product_variant.product_id").
			Where("product_variant.id = ?", ID).
	*/
	err := r.db.Model(&productVariants).
		ColumnExpr("product_variant.id, product_variant.product_id, product_variant.sku, product_variant.product_variant_name, product_variant.created_at, product_variant.updated_at").
		ColumnExpr("products.name AS product_name").
		Join("JOIN products ON products.id = product_variant.product_id").
		Select()
	if err != nil {
		productVariantRepositoryLogging.Printlog("GetAllproducts_Error", err.Error())
		return nil, err
	}

	return productVariants, nil
}

/*
GetProductVariantByID returns a product variant by the id from the product variant's table
*/
func (r *repository) GetProductVariantByID(ID uuid.UUID) (ProductVariant, error) {
	productVariant := ProductVariant{}

	err := r.db.Model(&productVariant).
		ColumnExpr("product_variant.id, product_variant.product_id, product_variant.sku, product_variant.product_variant_name, product_variant.created_at, product_variant.updated_at").
		ColumnExpr("products.name AS product_name").
		Join("JOIN products ON products.id = product_variant.product_id").
		Where("product_variant.id = ?", ID).
		Select()

	if err != nil {
		productVariantRepositoryLogging.Printlog("GetAllproducts_Error", err.Error())
		return ProductVariant{}, err
	}

	return productVariant, nil
}

/*
GetProductVariantsByName returns a product variant by the id from the product variant's table
*/
func (r *repository) GetProductVariantsByName(sku string) ([]ProductVariant, error) {
	productVariants := []ProductVariant{}
	err := r.db.Model(&productVariants).
		ColumnExpr("product_variant.id, product_variant.product_id, product_variant.sku, product_variant.product_variant_name, product_variant.created_at, product_variant.updated_at").
		ColumnExpr("products.name AS product_name").
		Join("JOIN products ON products.id = product_variant.product_id").
		Where("product_variant.sku like ?", "%"+sku+"%").
		Select()
	if err != nil {
		productVariantRepositoryLogging.Printlog("GetAllproducts_Error", err.Error())
		return nil, err
	}

	return productVariants, nil
}
