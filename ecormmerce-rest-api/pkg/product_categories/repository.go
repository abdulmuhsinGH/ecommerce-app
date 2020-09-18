package productcategories

import (
	"ecormmerce-app/ecormmerce-rest-api/pkg/logging"

	"github.com/go-pg/pg/v9"
)

/*
Repository provides productCategory repository operations
*/
type Repository interface {
	//AddproductCategory(*productCategory) bool
	GetAllProductCategories() ([]ProductCategory, error)
	AddProductCategory(*ProductCategory) (*ProductCategory, error)
	UpdateProductCategory(*ProductCategory) (*ProductCategory, error)
	DeleteProductCategory(*ProductCategory) error
	GetProductCategoryByID(int64) (ProductCategory, error)
	GetProductCategoriesByName(string) ([]ProductCategory, error)
}

type repository struct {
	db *pg.DB
}

var productCategoryRepositoryLogging logging.Logging

/*
NewRepository creates a ProductCategory repository with the necessary dependencies
*/
func NewRepository(db *pg.DB) Repository {
	productCategoryRepositoryLogging = logging.New("productCategory_repository: ")
	return &repository{db}
}

func (r *repository) UpdateProductCategory(productCategory *ProductCategory) (*ProductCategory, error) {
	_, err := r.db.Model(productCategory).Column("id", "name", "description", "created_at", "updated_by", "updated_at", "deleted_at").WherePK().Update()
	if err != nil {
		productCategoryRepositoryLogging.Printlog("UpdateproductCategory_Error", err.Error())
		return &ProductCategory{}, err
	}
	return productCategory, nil

}

/*
AddProductCategory finds productCategory or saves productCategory if not found to the productCategory's table
*/
func (r *repository) AddProductCategory(productCategory *ProductCategory) (*ProductCategory, error) {

	_, err := r.db.Model(productCategory).
		Returning("id").
		Insert()
	if err != nil {
		productCategoryRepositoryLogging.Printlog("AddproductCategory_Error", err.Error())
		return &ProductCategory{}, err
	}

	return productCategory, nil

}

/*
DeleteProductCategory saves productCategory to the productCategory's table
*/
func (r *repository) DeleteProductCategory(productCategory *ProductCategory) error {
	_, err := r.db.Model(productCategory).WherePK().Delete()
	if err != nil {
		productCategoryRepositoryLogging.Printlog("DeleteproductCategory_Error", err.Error())
		return err
	}
	return nil

}

/*
GetAllProductCategories returns all productCategorys from the productCategory's table
*/
func (r *repository) GetAllProductCategories() ([]ProductCategory, error) {
	productCategories := []ProductCategory{}
	err := r.db.Model(&productCategories).
		Column("id", "name", "description", "created_at", "updated_by", "updated_at", "deleted_at").
		Select()
	if err != nil {
		productCategoryRepositoryLogging.Printlog("GetAllproductCategorys_Error", err.Error())
		return nil, err
	}

	return productCategories, nil
}

/*
GetProductCategoryByID returns a productCategory by the id from the productCategory's table
*/
func (r *repository) GetProductCategoryByID(ID int64) (ProductCategory, error) {
	productCategory := ProductCategory{}

	err := r.db.Model(&productCategory).
		Column("id", "name", "description", "created_at", "updated_by", "updated_at", "deleted_at").
		Where("id = ?", ID).
		Select()

	if err != nil {
		productCategoryRepositoryLogging.Printlog("GetAllproductCategorys_Error", err.Error())
		return ProductCategory{}, err
	}

	return productCategory, nil
}

/*
GetProductCategoriesByName returns a productCategory by the id from the productCategory's table
*/
func (r *repository) GetProductCategoriesByName(name string) ([]ProductCategory, error) {
	productCategories := []ProductCategory{}
	err := r.db.Model(&productCategories).Where("name like ?", "%"+name+"%").
		Column("id", "name", "description", "created_at", "updated_by", "updated_at", "deleted_at").
		Select()
	if err != nil {
		productCategoryRepositoryLogging.Printlog("GetAllproductCategorys_Error", err.Error())
		return nil, err
	}

	return productCategories, nil
}
