package main

import (
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