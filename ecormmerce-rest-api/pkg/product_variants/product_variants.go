package productvariants

import (
	"time"

	"github.com/google/uuid"
)

//ProductVariant defines the properties of a produt type
type ProductVariant struct {
	ID                  uuid.UUID `json:"id"`
	ProductID           uuid.UUID `json:"product_id"`
	ProductName         string    `json:"product_name" pg:",discard_unknown_columns" sql:"-"`
	SKU                 string    `json:"sku"`
	ProductVariantValue string    `json:"product_variant_name"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	DeletedAt           time.Time `json:"deleted_at"`
}

//create product

//update product info

//read all products

//read one product

//delete a product
