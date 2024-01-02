package entity

import (
	"time"
)

// Album represents an album record.
type STA struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Property1 string `json:"property1"`
	Property2 string `json:"property2"`
	Property3 string `json:"property3"`
}
