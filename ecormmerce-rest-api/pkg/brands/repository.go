package brands

import (
	"ecormmerce-app/ecormmerce-rest-api/pkg/logging"

	"github.com/go-pg/pg/v9"
	"github.com/google/uuid"
)

/*
Repository provides Brand repository operations
*/
type Repository interface {
	//AddBrand(*Brand) bool
	GetAllBrands() ([]Brand, error)
	AddBrand(*Brand) (*Brand, error)
	UpdateBrand(*Brand) (*Brand, error)
	DeleteBrand(*Brand) error
	GetBrandByID(uuid.UUID) (Brand, error)
	GetBrandsByName(string) ([]Brand, error)
}

type repository struct {
	db *pg.DB
}

var brandRepositoryLogging logging.Logging

/*
NewRepository creates a Brand repository with the necessary dependencies
*/
func NewRepository(db *pg.DB) Repository {
	brandRepositoryLogging = logging.New("brand_repository: ")
	return &repository{db}
}

func (r *repository) UpdateBrand(brand *Brand) (*Brand, error) {
	_, err := r.db.Model(brand).Column("id", "name", "created_at", "updated_at", "deleted_at").WherePK().Update()
	if err != nil {
		brandRepositoryLogging.Printlog("UpdateBrand_Error", err.Error())
		return &Brand{}, err
	}
	return brand, nil

}

/*
FindOrAddBrand finds Brand or saves Brand if not found to the Brand's table
*/
func (r *repository) AddBrand(brand *Brand) (*Brand, error) {

	_, err := r.db.Model(brand).
		Returning("id").
		Insert()
	if err != nil {
		brandRepositoryLogging.Printlog("AddBrand_Error", err.Error())
		return &Brand{}, err
	}

	return brand, nil

}

/*
DeleteBrand saves Brand to the Brand's table
*/
func (r *repository) DeleteBrand(brand *Brand) error {
	_, err := r.db.Model(brand).WherePK().Delete()
	if err != nil {
		brandRepositoryLogging.Printlog("DeleteBrand_Error", err.Error())
		return err
	}
	return nil

}

/*
GetAllBrands returns all Brands from the Brand's table
*/
func (r *repository) GetAllBrands() ([]Brand, error) {
	var brands []Brand
	err := r.db.Model(&brands).
		Column("id", "name", "created_at", "updated_at", "deleted_at").
		Select()
	if err != nil {
		brandRepositoryLogging.Printlog("GetAllBrands_Error", err.Error())
		return nil, err
	}

	return brands, nil
}

/*
GetBrandByID returns a Brand by the id from the Brand's table
*/
func (r *repository) GetBrandByID(ID uuid.UUID) (Brand, error) {
	var brand Brand

	err := r.db.Model(&brand).
		Column("id", "name", "created_at", "updated_at", "deleted_at").
		Where("id = ?", ID).
		Select()

	if err != nil {
		brandRepositoryLogging.Printlog("GetAllBrands_Error", err.Error())
		return Brand{}, err
	}

	return brand, nil
}

/*
GetBrandsByName returns a Brand by the id from the Brand's table
*/
func (r *repository) GetBrandsByName(name string) ([]Brand, error) {
	var brands []Brand
	err := r.db.Model(&brands).Where("name like ?", "%"+name+"%").
		Column("id", "name", "created_at", "updated_at", "deleted_at").
		Select()
	if err != nil {
		brandRepositoryLogging.Printlog("GetAllBrands_Error", err.Error())
		return nil, err
	}

	return brands, nil
}
