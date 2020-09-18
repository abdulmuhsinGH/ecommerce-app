package products

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

//Product defines the properties of a produt type
type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Category    int       `json:"category"`
	Brand       int       `json:"brand"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedBy   int       `json:"updated_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

//create product

//update product info

//read all products

//read one product

//delete a product
