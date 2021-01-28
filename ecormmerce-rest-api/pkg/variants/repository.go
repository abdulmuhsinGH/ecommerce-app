package variants

import (
	"ecormmerce-app/ecormmerce-rest-api/pkg/logging"

	"github.com/go-pg/pg/v9"
	"github.com/google/uuid"
)

/*
Repository provides product repository operations
*/
type Repository interface {
	// AddVariant(*Variant) bool
	GetAllVariants() ([]Variant, error)
	GetAllVariantValues() ([]VariantValue, error)
	AddVariant(*Variant) (*Variant, error)
	AddVariantValue(*VariantValue) (*VariantValue, error)
	// UpdateVariantt(*Variant) (*Variant, error)
	// DeleteVariantt(*Variant) error
	GetVariantByID(uuid.UUID) (Variant, error)
	GetVariantsByName(string) ([]Variant, error)
	GetVariantValueByID(uuid.UUID) (VariantValue, error)
	GetVariantValuesByName(string) ([]VariantValue, error)
}

type repository struct {
	db *pg.DB
}

var variantRepositoryLogging logging.Logging

/*
NewRepository creates a variant repository with the necessary dependencies
*/
func NewRepository(db *pg.DB) Repository {
	variantRepositoryLogging = logging.New("variant_repository: ")
	return &repository{db}
}

/* func (r *repository) UpdateVariant(variant *Variant) (*Variant, error) {
	_, err := r.db.Model(variant).Column("id", "variant_name", "variant_desc", "updated_at").WherePK().Update()
	if err != nil {
		variantRepositoryLogging.Printlog("UpdateVariant_Error", err.Error())
		return &Variant{}, err
	}
	return variant, nil

} */

/*
AddVariant saves variant to the variants table
*/
func (r *repository) AddVariant(variant *Variant) (*Variant, error) {

	_, err := r.db.Model(variant).
		Returning("id").
		Insert()
	if err != nil {
		variantRepositoryLogging.Printlog("FindORAddVariant_Error", err.Error())
		return &Variant{}, err
	}

	return variant, nil

}

/*
AddVariantValue saves variantvalue to the variantvalues table
*/
func (r *repository) AddVariantValue(variantValue *VariantValue) (*VariantValue, error) {

	_, err := r.db.Model(variantValue).
		Returning("id").
		Insert()
	if err != nil {
		variantRepositoryLogging.Printlog("FindORAddVariantValue_Error", err.Error())
		return &VariantValue{}, err
	}

	return variantValue, nil

}

/*
DeleteVariant saves vairant to the variants table
*/
// func (r *repository) DeleteVariantt(variant *Variant) error {
// 	_, err := r.db.Model(variant).WherePK().Delete()
// 	if err != nil {
// 		variantRepositoryLogging.Printlog("DeleteVariantt_Error", err.Error())
// 		return err
// 	}
// 	return nil

// }

/*
GetAllVariants returns all variants from the variantss table
*/
func (r *repository) GetAllVariants() ([]Variant, error) {
	variants := []Variant{}
	err := r.db.Model(&variants).
		Column("id", "variant_name", "variant_desc", "updated_by", "created_at", "updated_at").
		Select()
	if err != nil {
		variantRepositoryLogging.Printlog("GetAllvariants_Error", err.Error())
		return nil, err
	}

	return variants, nil
}

/*
GetAllVariantValues returns all variantvalues from the variant_value table
*/
func (r *repository) GetAllVariantValues() ([]VariantValue, error) {
	variantValues := []VariantValue{}
	err := r.db.Model(&variantValues).
		Column("id", "variant_id", "variant_value_name", "updated_by", "created_at", "updated_at").
		Select()
	if err != nil {
		variantRepositoryLogging.Printlog("GetAllVariantValuess_Error", err.Error())
		return nil, err
	}

	return variantValues, nil
}

/*
GetVariantByID returns a variant by the id from the variants table
*/
func (r *repository) GetVariantByID(ID uuid.UUID) (Variant, error) {
	variant := Variant{}

	err := r.db.Model(&variant).
		Column("id", "variant_name", "variant_desc", "updated_by", "created_at", "updated_at").
		Where("id = ?", ID).
		Select()

	if err != nil {
		variantRepositoryLogging.Printlog("GetVariantByID_Error", err.Error())
		return Variant{}, err
	}

	return variant, nil
}

/*
GetVariantValueByID returns a variantValue by the id from the variant_value table
*/
func (r *repository) GetVariantValueByID(ID uuid.UUID) (VariantValue, error) {
	variantValue := VariantValue{}

	err := r.db.Model(&variantValue).
		Column("id", "variant_id", "variant_value_name", "updated_by", "created_at", "updated_at").
		Where("id = ?", ID).
		Select()

	if err != nil {
		variantRepositoryLogging.Printlog("GetAllVariantValueByID_Error", err.Error())
		return VariantValue{}, err
	}

	return variantValue, nil
}

/*
GetVariantsByName returns a variant by the id from the variants table
*/
func (r *repository) GetVariantsByName(name string) ([]Variant, error) {
	variants := []Variant{}
	err := r.db.Model(&variants).Where("variant_name like ?", "%"+name+"%").
		Column("id", "variant_name", "variant_desc", "updated_by", "created_at", "updated_at").
		Select()
	if err != nil {
		variantRepositoryLogging.Printlog("GetAllVariants", err.Error())
		return nil, err
	}

	return variants, nil
}

/*
GetVariantValuesByName returns a variant_value_name  by the id from the variant_value table
*/
func (r *repository) GetVariantValuesByName(name string) ([]VariantValue, error) {
	variantValues := []VariantValue{}
	err := r.db.Model(&variantValues).Where("variant_value_name like ?", "%"+name+"%").
		Column("id", "variant_value_name", "created_at", "updated_by", "updated_at").
		Select()
	if err != nil {
		variantRepositoryLogging.Printlog("GetVariantValuesByName", err.Error())
		return nil, err
	}

	return variantValues, nil
}
