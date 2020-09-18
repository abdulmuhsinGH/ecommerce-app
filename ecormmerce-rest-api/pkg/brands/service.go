package brands

import (
	"ecormmerce-app/ecormmerce-rest-api/pkg/logging"
	"errors"
	"time"
)

// Service provides ProductBrand adding operations.
type Service interface {
	AddBrand(*ProductBrand) error
	GetAllBrands() ([]ProductBrand, error)
	UpdateBrand(ProductBrand *ProductBrand) error
	DeleteBrand(ProductBrand *ProductBrand) error
	GetBrandByID(ID int) (ProductBrand, error)
	GetBrandsByName(name string) ([]ProductBrand, error)
}

type service struct {
	brandRepository Repository
}

var brandServiceLogging logging.Logging

/*
NewService creates a Brands service with the necessary dependencies
*/
func NewService(r Repository) Service {
	brandServiceLogging = logging.New("brand_service:")
	return &service{r}
}

/*
AddBrand creates a new ProductBrand
*/
func (s *service) AddBrand(brand *ProductBrand) error {

	brand, err := s.brandRepository.AddBrand(brand)
	if err != nil {
		return err
	}
	return nil

}

/*
UpdateBrand creates a new ProductBrand
*/
func (s *service) UpdateBrand(brand *ProductBrand) error {
	var err error

	brand.UpdatedAt = time.Now()
	brand, err = s.brandRepository.UpdateBrand(brand)
	if err != nil {
		return err
	}
	return nil

}

/*
DeleteBrand creates a new ProductBrand
*/
func (s *service) DeleteBrand(brand *ProductBrand) error {
	err := s.brandRepository.DeleteBrand(brand)
	if err != nil {
		return errors.New("not deleted")
	}
	return nil

}

/*
GetAllBrands gets all Brands
*/
func (s *service) GetAllBrands() ([]ProductBrand, error) {
	brands, err := s.brandRepository.GetAllBrands()
	if err != nil {
		return nil, err
	}
	return brands, nil
}

/*
GetBrandByID gets all Brands
*/
func (s *service) GetBrandByID(ID int) (ProductBrand, error) {
	brand, err := s.brandRepository.GetBrandByID(ID)
	if err != nil {
		return ProductBrand{}, err
	}
	return brand, nil
}

/*
GetBrandsByName gets all Brands
*/
func (s *service) GetBrandsByName(name string) ([]ProductBrand, error) {
	brands, err := s.brandRepository.GetBrandsByName(name)
	if err != nil {
		return nil, err
	}
	return brands, nil
}
