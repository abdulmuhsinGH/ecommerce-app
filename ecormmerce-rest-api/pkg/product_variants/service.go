package productvariants

import (
	"ecormmerce-app/ecormmerce-rest-api/pkg/logging"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Service provides productVariant adding operations.
type Service interface {
	AddProductVariant(*ProductVariant) error
	GetAllProductVariants() ([]ProductVariant, error)
	UpdateProductVariant(productVariant *ProductVariant) error
	DeleteProductVariant(productVariant *ProductVariant) error
	GetProductVariantByID(ID uuid.UUID) (ProductVariant, error)
	GetProductVariantsByName(name string) ([]ProductVariant, error)
}

type service struct {
	productVariantRepository Repository
}

var productVariantServiceLogging logging.Logging

/*
NewService creates a products service with the necessary dependencies
*/
func NewService(r Repository) Service {
	productVariantServiceLogging = logging.New("product_service:")
	return &service{r}
}

/*
AddProduct creates a new productVariant
*/
func (s *service) AddProductVariant(productVariant *ProductVariant) error {

	productVariant, err := s.productVariantRepository.AddProductVariant(productVariant)
	if err != nil {
		return err
	}
	return nil

}

/*
UpdateProduct creates a new productVariant
*/
func (s *service) UpdateProductVariant(productVariant *ProductVariant) error {
	productVariant.UpdatedAt = time.Now().UTC()
	_, err := s.productVariantRepository.UpdateProductVariant(productVariant)
	if err != nil {
		return err
	}
	return nil

}

/*
DeleteProduct creates a new productVariant
*/
func (s *service) DeleteProductVariant(productVariant *ProductVariant) error {
	err := s.productVariantRepository.DeleteProductVariant(productVariant)
	if err != nil {
		return errors.New("not deleted")
	}
	return nil

}

/*
GetAllProducts gets all products
*/
func (s *service) GetAllProductVariants() ([]ProductVariant, error) {
	products, err := s.productVariantRepository.GetAllProductVariants()
	if err != nil {
		return nil, err
	}
	return products, nil
}

/*
GetProductByID gets all products
*/
func (s *service) GetProductVariantByID(ID uuid.UUID) (ProductVariant, error) {
	productVariant, err := s.productVariantRepository.GetProductVariantByID(ID)
	if err != nil {
		return ProductVariant{}, err
	}
	return productVariant, nil
}

/*
GetProductsByName gets all products
*/
func (s *service) GetProductVariantsByName(name string) ([]ProductVariant, error) {
	products, err := s.productVariantRepository.GetProductVariantsByName(name)
	if err != nil {
		return nil, err
	}
	return products, nil
}
