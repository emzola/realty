package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/emzola/realty/internal/data"
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
		ID: id,
		CreatedAt: time.Now(),
		Title: "6007 Applegate Lane",
		Description: "Don't let him know she liked them best, For this must ever be A secret, kept from all the children she knew she had a bone in his throat,' said the Footman, 'and that for the moment they saw the White Rabbit as he spoke, and added 'It isn't mine,' said the King.",
		City: "Moscow",
		Location: "65 Tverskaya street",
		Latitude: 1.225,
		Longitude: 3.664,
		Type: "For Sale",
		Category: "Villa",
		Features: map[string]int32{
			"Bedrooms": 1,
			"Bathrooms": 1,
			"Floors": 3,
			"SquareMetres": 83,
		},
		Price: 200000,
		Currency: "USD",
		Nearby: map[string]string{
			"Hospital": "7km",
			"Busstop": "12km",
		},
		Amenities: []string{"Parking", "Laundry Room"},
		Version: 1,
	}

	err = app.writeJSON(w, http.StatusOK, envelop{"property": property}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func(app *application) createPropertyHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
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
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}