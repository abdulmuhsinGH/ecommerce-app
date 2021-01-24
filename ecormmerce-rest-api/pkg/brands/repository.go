package brands

import (
	"ecormmerce-app/ecormmerce-rest-api/pkg/logging"

	"github.com/go-pg/pg/v9"
)

/*
Repository provides ProductBrand repository operations
*/
type Repository interface {
	//AddBrand(*ProductBrand) bool
	GetAllBrands() ([]ProductBrand, error)
	AddBrand(*ProductBrand) (*ProductBrand, error)
	UpdateBrand(*ProductBrand) (*ProductBrand, error)
	DeleteBrand(*ProductBrand) error
	GetBrandByID(int) (ProductBrand, error)
	GetBrandsByName(string) ([]ProductBrand, error)
}

type repository struct {
	db *pg.DB
}

var brandRepositoryLogging logging.Logging

/*
NewRepository creates a ProductBrand repository with the necessary dependencies
*/
func NewRepository(db *pg.DB) Repository {
	brandRepositoryLogging = logging.New("brand_repository: ")
	return &repository{db}
}

func (r *repository) UpdateBrand(brand *ProductBrand) (*ProductBrand, error) {
	_, err := r.db.Model(brand).Column("id", "name", "updated_at").WherePK().Update()
	if err != nil {
		brandRepositoryLogging.Printlog("UpdateBrand_Error", err.Error())
		return &ProductBrand{}, err
	}
	return brand, nil

}

/*
FindOrAddBrand finds ProductBrand or saves ProductBrand if not found to the ProductBrand's table
*/
func (r *repository) AddBrand(brand *ProductBrand) (*ProductBrand, error) {

	_, err := r.db.Model(brand).
		Returning("id").
		Insert()
	if err != nil {
		brandRepositoryLogging.Printlog("AddBrand_Error", err.Error())
		return &ProductBrand{}, err
	}

	return brand, nil

}

/*
DeleteBrand saves ProductBrand to the ProductBrand's table
*/
func (r *repository) DeleteBrand(brand *ProductBrand) error {
	_, err := r.db.Model(brand).WherePK().Delete()
	if err != nil {
		brandRepositoryLogging.Printlog("DeleteBrand_Error", err.Error())
		return err
	}
	return nil

}

/*
GetAllBrands returns all Brands from the ProductBrand's table
*/
func (r *repository) GetAllBrands() ([]ProductBrand, error) {
	brands := []ProductBrand{}
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
GetBrandByID returns a ProductBrand by the id from the ProductBrand's table
*/
func (r *repository) GetBrandByID(ID int) (ProductBrand, error) {
	brand := ProductBrand{}

	err := r.db.Model(&brand).
		Column("id", "name", "created_at", "updated_at", "deleted_at").
		Where("id = ?", ID).
		Select()

	if err != nil {
		brandRepositoryLogging.Printlog("GetAllBrands_Error", err.Error())
		return ProductBrand{}, err
	}

	return brand, nil
}

/*
GetBrandsByName returns a ProductBrand by the id from the ProductBrand's table
*/
func (r *repository) GetBrandsByName(name string) ([]ProductBrand, error) {
	brands := []ProductBrand{}
	err := r.db.Model(&brands).Where("name like ?", "%"+name+"%").
		Column("id", "name", "created_at", "updated_at", "deleted_at").
		Select()
	if err != nil {
		brandRepositoryLogging.Printlog("GetAllBrands_Error", err.Error())
		return nil, err
	}

	return brands, nil
}
