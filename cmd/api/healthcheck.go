package main

import (
	"net/http"
)

// healthcheckHandler shows application information.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	health := envelop{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version": version,
		},
	}

 err := app.writeJSON(w, http.StatusOK, health, nil)
 if err != nil {
	app.logger.Println(err)
	http.Error(w, "the server encountered an error and could not process your request", http.StatusInternalServerError)
	return
 }
}