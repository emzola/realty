package data

import "time"

// Property contains information about a property
type Property struct {
	ID        int64	`json:"id"`
	CreatedAt time.Time	`json:"-"`
	Title string	`json:"title"`
	Description string	`json:"description"`
	City string	`json:"city"`
	Location string	`json:"location"`
	Latitude float64	`json:"latitude,omitempty"`
	Longitude float64	`json:"longitude,omitempty"`
	Type string	`json:"type,omitempty"`
	Category string	`json:"category,omitempty"`
	Features map[string]int32	`json:"features,omitempty"`
	Price float64	`json:"price"`
	Currency string	`json:"currency"`
	Nearby map[string]string	`json:"nearby,omitempty"`
	Amenities []string	`json:"amenities,omitempty"`
	Version int32	`json:"version"`
}