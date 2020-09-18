package productcategories

import (
	"ecormmerce-app/ecormmerce-rest-api/pkg/logging"
	"errors"
	"time"
)

// Service provides productCategory adding operations.
type Service interface {
	AddProductCategory(*ProductCategory) error
	GetAllProductCategories() ([]ProductCategory, error)
	UpdateProductCategory(productCategory *ProductCategory) error
	DeleteProductCategory(productCategory *ProductCategory) error
	GetProductCategoryByID(ID int64) (ProductCategory, error)
	GetProductCategoriesByName(name string) ([]ProductCategory, error)
}

type service struct {
	productCategoryRepository Repository
}

var productCategoryServiceLogging logging.Logging

/*
NewService creates a productCategorys service with the necessary dependencies
*/
func NewService(r Repository) Service {
	productCategoryServiceLogging = logging.New("productCategory_service:")
	return &service{r}
}

/*
AddproductCategory creates a new productCategory
*/
func (s *service) AddProductCategory(productCategory *ProductCategory) error {

	productCategory, err := s.productCategoryRepository.AddProductCategory(productCategory)
	if err != nil {
		return err
	}
	return nil

}

/*
UpdateproductCategory creates a new productCategory
*/
func (s *service) UpdateProductCategory(productCategory *ProductCategory) error {
	var err error

	productCategory.UpdatedAt = time.Now()
	productCategory, err = s.productCategoryRepository.UpdateProductCategory(productCategory)
	if err != nil {
		return err
	}
	return nil

}

/*
DeleteproductCategory creates a new productCategory
*/
func (s *service) DeleteProductCategory(productCategory *ProductCategory) error {
	err := s.productCategoryRepository.DeleteProductCategory(productCategory)
	if err != nil {
		return errors.New("not deleted")
	}
	return nil

}

/*
GetAllProductCategories gets all productCategorys
*/
func (s *service) GetAllProductCategories() ([]ProductCategory, error) {
	productCategorys, err := s.productCategoryRepository.GetAllProductCategories()
	if err != nil {
		return nil, err
	}
	return productCategorys, nil
}

/*
GetproductCategoryByID gets all productCategorys
*/
func (s *service) GetProductCategoryByID(ID int64) (productCategory ProductCategory, err error) {
	productCategory, err = s.productCategoryRepository.GetProductCategoryByID(ID)
	if err != nil {
		return ProductCategory{}, err
	}
	return productCategory, nil
}

/*
GetProductCategoriesByName gets all productCategories
*/
func (s *service) GetProductCategoriesByName(name string) ([]ProductCategory, error) {
	productCategories, err := s.productCategoryRepository.GetProductCategoriesByName(name)
	if err != nil {
		return nil, err
	}
	return productCategories, nil
}
