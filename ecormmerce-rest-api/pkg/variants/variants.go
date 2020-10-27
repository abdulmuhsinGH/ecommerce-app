package variants

import (
	"time"

	"github.com/google/uuid"
)

//Variant defines the properties of a produt type
type Variant struct {
	ID          uuid.UUID `json:"id"`
	VariantName string    `json:"variant_name"`
	VariantDesc string    `json:"variant_desc"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedBy   uuid.UUID `json:"updated_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

type VariantValue struct {
	ID          uuid.UUID `json:"id"`
	VariantId   uuid.UUID `json:"variant_id"`
	VariantName string    `json:"variant_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

//create variant

//read all variants

//read one variant