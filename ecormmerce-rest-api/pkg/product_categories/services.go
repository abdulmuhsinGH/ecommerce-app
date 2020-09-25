package productcategories

import (
	"ecormmerce-app/ecormmerce-rest-api/pkg/logging"
	"errors"
	"time"
)

// Service provides product category adding operations.
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
NewService creates a product category service with the necessary dependencies
*/
func NewService(r Repository) Service {
	productCategoryServiceLogging = logging.New("productCategory_service:")
	return &service{r}
}

/*
AddProductCategory creates a new product category
*/
func (s *service) AddProductCategory(productCategory *ProductCategory) error {

	productCategory, err := s.productCategoryRepository.AddProductCategory(productCategory)
	if err != nil {
		return err
	}
	return nil

}

/*
UpdateProductCategory creates a new product category
*/
func (s *service) UpdateProductCategory(productCategory *ProductCategory) error {
	productCategory.UpdatedAt = time.Now().UTC()
	_, err := s.productCategoryRepository.UpdateProductCategory(productCategory)
	if err != nil {
		return err
	}
	return nil

}

/*
DeleteProductCategory creates a new product category
*/
func (s *service) DeleteProductCategory(productCategory *ProductCategory) error {
	err := s.productCategoryRepository.DeleteProductCategory(productCategory)
	if err != nil {
		return errors.New("not deleted")
	}
	return nil

}

/*
GetAllProductCategories gets all product categorys
*/
func (s *service) GetAllProductCategories() ([]ProductCategory, error) {
	productCategorys, err := s.productCategoryRepository.GetAllProductCategories()
	if err != nil {
		return nil, err
	}
	return productCategorys, nil
}

/*
GetProductCategoryByID gets a product category using the ID
*/
func (s *service) GetProductCategoryByID(ID int64) (productCategory ProductCategory, err error) {
	productCategory, err = s.productCategoryRepository.GetProductCategoryByID(ID)
	if err != nil {
		return ProductCategory{}, err
	}
	return productCategory, nil
}

/*
GetProductCategoriesByName gets a product category using the name
*/
func (s *service) GetProductCategoriesByName(name string) ([]ProductCategory, error) {
	productCategories, err := s.productCategoryRepository.GetProductCategoriesByName(name)
	if err != nil {
		return nil, err
	}
	return productCategories, nil
}
