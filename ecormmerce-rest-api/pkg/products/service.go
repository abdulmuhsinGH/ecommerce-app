package products

import (
	"ecormmerce-app/ecormmerce-rest-api/pkg/logging"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Service provides product adding operations.
type Service interface {
	AddProduct(*Product) error
	GetAllProducts() ([]Product, error)
	UpdateProduct(product *Product) error
	DeleteProduct(product *Product) error
	GetProductByID(ID uuid.UUID) (Product, error)
	GetProductsByName(name string) ([]Product, error)
}

type service struct {
	productRepository Repository
}

var productServiceLogging logging.Logging

/*
NewService creates a products service with the necessary dependencies
*/
func NewService(r Repository) Service {
	productServiceLogging = logging.New("product_service:")
	return &service{r}
}

/*
AddProduct creates a new product
*/
func (s *service) AddProduct(product *Product) error {

	product, err := s.productRepository.AddProduct(product)
	if err != nil {
		return err
	}
	return nil

}

/*
UpdateProduct creates a new product
*/
func (s *service) UpdateProduct(product *Product) error {
	product.UpdatedAt = time.Now().UTC()
	_, err := s.productRepository.UpdateProduct(product)
	if err != nil {
		return err
	}
	return nil

}

/*
DeleteProduct creates a new product
*/
func (s *service) DeleteProduct(product *Product) error {
	err := s.productRepository.DeleteProduct(product)
	if err != nil {
		return errors.New("not deleted")
	}
	return nil

}

/*
GetAllProducts gets all products
*/
func (s *service) GetAllProducts() ([]Product, error) {
	products, err := s.productRepository.GetAllProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

/*
GetProductByID gets all products
*/
func (s *service) GetProductByID(ID uuid.UUID) (Product, error) {
	product, err := s.productRepository.GetProductByID(ID)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

/*
GetProductsByName gets all products
*/
func (s *service) GetProductsByName(name string) ([]Product, error) {
	products, err := s.productRepository.GetProductsByName(name)
	if err != nil {
		return nil, err
	}
	return products, nil
}
