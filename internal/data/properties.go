package data

import (
	"time"

	"github.com/emzola/realty/internal/validator"
)

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

// ValidateProperty validates a property based on set validation criteria.
func ValidateProperty(v *validator.Validator, property *Property) {
	v.Check(property.Title != "", "title", "must be provided")
	v.Check(len(property.Title) <= 500, "title", "must not be more than 500 bytes long")
	v.Check(property.Description != "", "description", "must be provided")
	v.Check(property.City != "", "city", "must be provided")
	v.Check(property.Location != "", "location", "must be provided")
	v.Check(property.Type != "", "type", "must be a provided")
	v.Check(property.Category != "", "category", "must be provided")
	v.Check(property.Features != nil, "features", "must be provided")
	v.Check(len(property.Features) >= 1, "features", "must contain at least 1 feature")
	v.Check(len(property.Features) <= 20, "features", "must not contain more than 20 features")
	v.Check(property.Price != 0, "price", "must be provided")
	v.Check(property.Price > 0, "price", "must be a positive number")
	v.Check(property.Currency != "", "currency", "must be provided")
	v.Check(property.Nearby != nil, "nearby", "must be provided")
	v.Check(len(property.Nearby) >= 1, "nearby", "must contain at least 1 facility")
	v.Check(len(property.Nearby) <= 10, "nearby", "must not contain more than 10 facilities")
	v.Check(validator.Unique(property.Amenities), "amenities", "must not contain duplicate values")
}