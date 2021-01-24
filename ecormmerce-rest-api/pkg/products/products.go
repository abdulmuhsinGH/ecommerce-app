package products

import (
	"time"

	"github.com/google/uuid"
)

//Product defines the properties of a produt type
type Product struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Category    int64     `json:"category"`
	Brand       int64     `json:"brand"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedBy   uuid.UUID `json:"updated_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

//create product

//update product info

//read all products

//read one product

//delete a product
