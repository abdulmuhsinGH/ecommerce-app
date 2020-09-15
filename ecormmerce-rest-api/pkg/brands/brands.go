package brands

import (
	"time"
)

//Brand defines the properties of a produt type
type Brand struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

//create brand

//update brand info

//read all brands

//read one brand

//delete a brand
