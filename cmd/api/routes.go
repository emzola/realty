package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// routes returns a router of all api endpoints.
func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/properties/:id", app.showPropertyHandler)
	router.HandlerFunc(http.MethodPost, "/v1/account/properties/create", app.createPropertyHandler)

	return router
}
