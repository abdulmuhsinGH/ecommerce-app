package product_categories

import (
	"time"
)

//ProductCategory defines the properties of a produt type
type ProductCategory struct {
	ID          int64       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedBy   int       `json:"updated_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
