package brands

import (
	"ecormmerce-app/ecormmerce-rest-api/pkg/logging"
	"errors"
	"time"
)

// Service provides Brand adding operations.
type Service interface {
	AddBrand(*Brand) error
	GetAllBrands() ([]Brand, error)
	UpdateBrand(Brand *Brand) error
	DeleteBrand(Brand *Brand) error
	GetBrandByID(ID int) (Brand, error)
	GetBrandsByName(name string) ([]Brand, error)
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
AddBrand creates a new Brand
*/
func (s *service) AddBrand(brand *Brand) error {

	brand, err := s.brandRepository.AddBrand(brand)
	if err != nil {
		return err
	}
	return nil

}

/*
UpdateBrand creates a new Brand
*/
func (s *service) UpdateBrand(brand *Brand) error {
	var err error

	brand.UpdatedAt = time.Now()
	brand, err = s.brandRepository.UpdateBrand(brand)
	if err != nil {
		return err
	}
	return nil

}

/*
DeleteBrand creates a new Brand
*/
func (s *service) DeleteBrand(brand *Brand) error {
	err := s.brandRepository.DeleteBrand(brand)
	if err != nil {
		return errors.New("not deleted")
	}
	return nil

}

/*
GetAllBrands gets all Brands
*/
func (s *service) GetAllBrands() ([]Brand, error) {
	brands, err := s.brandRepository.GetAllBrands()
	if err != nil {
		return nil, err
	}
	return brands, nil
}

/*
GetBrandByID gets all Brands
*/
func (s *service) GetBrandByID(ID int) (Brand, error) {
	brand, err := s.brandRepository.GetBrandByID(ID)
	if err != nil {
		return Brand{}, err
	}
	return brand, nil
}

/*
GetBrandsByName gets all Brands
*/
func (s *service) GetBrandsByName(name string) ([]Brand, error) {
	brands, err := s.brandRepository.GetBrandsByName(name)
	if err != nil {
		return nil, err
	}
	return brands, nil
}
