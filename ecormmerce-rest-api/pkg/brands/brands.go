package brands

import (
	"time"
)

//ProductBrand defines the properties of a produt type
type ProductBrand struct {
	tableName struct{} `pg:"product_brands"`
	ID        int64    `json:"id"`
	Name      string   `json:"name"`
	//Description string    `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

//create brand

//update brand info

//read all brands

//read one brand

//delete a brand
