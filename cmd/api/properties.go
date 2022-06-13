package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/emzola/realty/internal/data"
	"github.com/emzola/realty/internal/validator"
)

// showPropertyHandler shows property details.
func (app *application) showPropertyHandler(w http.ResponseWriter, r *http.Request) {
	// extract ID param
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	// instantiate sample property struct
	property := data.Property{
		ID:          id,
		CreatedAt:   time.Now(),
		Title:       "6007 Applegate Lane",
		Description: "Don't let him know she liked them best, For this must ever be A secret, kept from all the children she knew she had a bone in his throat,' said the Footman, 'and that for the moment they saw the White Rabbit as he spoke, and added 'It isn't mine,' said the King.",
		City:        "Moscow",
		Location:    "65 Tverskaya street",
		Latitude:    1.225,
		Longitude:   3.664,
		Type:        []string{"For Sale"},
		Category:    []string{"Villa"},
		Features: data.Features{
			"Bedrooms":     1,
			"Bathrooms":    1,
			"Floors":       3,
			"SquareMetres": 83,
		},
		Price:    200000,
		Currency: []string{"USD"},
		Nearby: data.Nearby{
			"Hospital": "7km",
			"Busstop":  "12km",
		},
		Amenities: []string{"Parking", "Laundry Room"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelop{"property": property}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// createPropertyHandler creates a property.
func (app *application) createPropertyHandler(w http.ResponseWriter, r *http.Request) {
	// Decode JSON into this input struct instead of directly on the property struct.
	// That way, the client does not have to provide ID and Version fields
	var input struct {
		Title       string            `json:"title"`
		Description string            `json:"description"`
		City        string            `json:"city"`
		Location    string            `json:"location"`
		Latitude    float64           `json:"latitude,omitempty"`
		Longitude   float64           `json:"longitude,omitempty"`
		Type        []string          `json:"type,omitempty"`
		Category    []string          `json:"category,omitempty"`
		Features   	data.Features  		`json:"features,omitempty"`
		Price       float64           `json:"price"`
		Currency    []string          `json:"currency"`
		Nearby      data.Nearby				`json:"nearby,omitempty"`
		Amenities   []string          `json:"amenities,omitempty"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Copy values from the input struct into a new property struct
	property := &data.Property{
		Title:       input.Title,
		Description: input.Description,
		City:        input.City,
		Location:    input.Location,
		Latitude:    input.Latitude,
		Longitude:   input.Longitude,
		Type:        input.Type,
		Category:    input.Category,
		Features:    input.Features,
		Price:       input.Price,
		Currency:    input.Currency,
		Nearby:      input.Nearby,
		Amenities:   input.Amenities,
	}

	v := validator.New()
	if data.ValidateProperty(v, property); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Insert data into properties DB table
	err = app.models.Properties.Insert(property)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// Set location header
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/properties/%d", property.ID))

	err = app.writeJSON(w, http.StatusCreated, envelop{"property": property}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
