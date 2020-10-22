package variants

import (
	"time"

	"github.com/google/uuid"
)

//Product defines the properties of a produt type
type Variant struct {
	ID           uuid.UUID `json:"id"`
	Variant_name string    `json:"variant_name"`
	Variant_desc string    `json:"variant_desc"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedBy    uuid.UUID `json:"updated_by"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

//create variant

//read all variants

//read one variant
