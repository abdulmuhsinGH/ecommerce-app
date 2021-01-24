package variants

import (
	"ecormmerce-app/ecormmerce-rest-api/pkg/logging"

	"github.com/google/uuid"
)

// Service provides product adding operations.
type Service interface {
	AddVariant(*Variant) error
	AddVariantValue(*VariantValue) error
	GetAllVariants() ([]Variant, error)
	GetAllVariantValues() ([]VariantValue, error)
	// UpdateVariant(variant *Variant) error
	// DeleteVariant(variant *Variant) error
	GetVariantByID(ID uuid.UUID) (Variant, error)
	GetVariantValueByID(ID uuid.UUID) (VariantValue, error)
	GetVariantsByName(name string) ([]Variant, error)
	GetVariantValuesByName(name string) ([]VariantValue, error)
}

type service struct {
	variantRepository Repository
}

var variantServiceLogging logging.Logging

/*
NewService creates a variant service with the necessary dependencies
*/
func NewService(r Repository) Service {
	variantServiceLogging = logging.New("variant_service:")
	return &service{r}
}

/*
AddVariant creates a new variant
*/
func (s *service) AddVariant(variant *Variant) error {

	variant, err := s.variantRepository.AddVariant(variant)
	if err != nil {
		return err
	}
	return nil

}

/*
AddVariantValue creates a new variantvalue
*/
func (s *service) AddVariantValue(variantValue *VariantValue) error {

	variantValue, err := s.variantRepository.AddVariantValue(variantValue)
	if err != nil {
		return err
	}
	return nil

}

/*
UpdateVariant creates a new variant
*/
// func (s *service) UpdateVariant(variant *Variant) error {
// 	variant.UpdatedAt = time.Now().UTC()
// 	_, err := s.variantRepository.UpdateVariant(variant)
// 	if err != nil {
// 		return err
// 	}
// 	return nil

// }

// /*
// DeleteVariant creates a new variant
// */
// func (s *service) DeleteVariant(variant *Variant) error {
// 	err := s.variantRepository.DeleteVariant(variant)
// 	if err != nil {
// 		return errors.New("not deleted")
// 	}
// 	return nil

// }

/*
GetAllVariants gets all variants
*/
func (s *service) GetAllVariants() ([]Variant, error) {
	variants, err := s.variantRepository.GetAllVariants()
	if err != nil {
		return nil, err
	}
	return variants, nil
}

/*
GetVariantByID gets all variants
*/
func (s *service) GetVariantByID(ID uuid.UUID) (Variant, error) {
	variant, err := s.variantRepository.GetVariantByID(ID)
	if err != nil {
		return Variant{}, err
	}
	return variant, nil
}

/*
GetVariantsByName gets all variant with the 'name'
*/
func (s *service) GetVariantsByName(name string) ([]Variant, error) {
	variants, err := s.variantRepository.GetVariantsByName(name)
	if err != nil {
		return nil, err
	}
	return variants, nil
}

/*
GetAllVariantValues gets all variantvaluess
*/
func (s *service) GetAllVariantValues() ([]VariantValue, error) {
	variantValues, err := s.variantRepository.GetAllVariantValues()
	if err != nil {
		return nil, err
	}
	return variantValues, nil
}

/*
GetVariantValueByID gets all variantvalues
*/
func (s *service) GetVariantValueByID(ID uuid.UUID) (VariantValue, error) {
	variantValue, err := s.variantRepository.GetVariantValueByID(ID)
	if err != nil {
		return VariantValue{}, err
	}
	return variantValue, nil
}

/*
GetVariantValuesByName gets all variant values  with the 'name'
*/
func (s *service) GetVariantValuesByName(name string) ([]VariantValue, error) {
	variantValues, err := s.variantRepository.GetVariantValuesByName(name)
	if err != nil {
		return nil, err
	}
	return variantValues, nil
}
