package main

import (
	"errors"
	"fmt"
	"net/http"

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

	property, err := app.models.Properties.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
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

	// Validate the property record, sending the client a 422 Unprocessable Entity
	// response if any checks fail
	v := validator.New()
	if data.ValidateProperty(v, property); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Insert property record into properties DB table
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

// updatePropertyHandler updates a property.
func(app *application) updatePropertyHandler(w http.ResponseWriter, r *http.Request) {
	// extract ID param
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	// Fetch corresponding property record from the database
	property, err := app.models.Properties.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// Decode JSON into this input struct instead of directly on the property struct.
	// That way, the client does not have to provide ID and Version fields
	var input struct {
		Title       *string            `json:"title"`
		Description *string            `json:"description"`
		City        *string            `json:"city"`
		Location    *string            `json:"location"`
		Latitude    *float64           `json:"latitude,omitempty"`
		Longitude   *float64           `json:"longitude,omitempty"`
		Type        []string          `json:"type,omitempty"`
		Category    []string          `json:"category,omitempty"`
		Features   	data.Features  		`json:"features,omitempty"`
		Price       *float64           `json:"price"`
		Currency    []string          `json:"currency"`
		Nearby      data.Nearby				`json:"nearby,omitempty"`
		Amenities   []string          `json:"amenities,omitempty"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// copy data acrosss from input struct to property record
	if input.Title != nil {
		property.Title = *input.Title
	}	
	if input.Description != nil {
		property.Description = *input.Description
	}
	if input.City != nil {
		property.City = *input.City
	}
	if input.Location != nil {
		property.Location = *input.Location
	}
	if input.Latitude != nil {
		property.Latitude = *input.Latitude
	}
	if input.Type != nil {
		property.Type = input.Type
	}
	if input.Category != nil {
		property.Category = input.Category
	}
	if input.Features != nil {
		property.Features = input.Features
	}
	if input.Price != nil {
		property.Price = *input.Price
	}
	if input.Currency != nil {
		property.Currency = input.Currency
	}
	if input.Nearby != nil {
		property.Nearby = input.Nearby
	}
	if input.Amenities != nil {
		property.Amenities = input.Amenities
	}

	// Validate the updated property record, sending the client a 422 Unprocessable Entity
	// response if any checks fail
	v := validator.New()
	if data.ValidateProperty(v, property); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Pass the updated property record to the Update() method to update the database
	err = app.models.Properties.Update(property)
	if err != nil {  
		app.serverErrorResponse(w, r, err)
		return
	}

	// Write the updated property record in a JSON response
	err = app.writeJSON(w, http.StatusOK, envelop{"property": property}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
// deletePropertyHandler deletes a property.
func (app *application) deletePropertyHandler(w http.ResponseWriter, r *http.Request) {
	// extract ID param
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Properties.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelop{"message": "property successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
