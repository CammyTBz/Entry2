// Filename: cmd/api/routes
package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	// Create a new HttpRouter router instance
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/entries", app.createEntryHandler) //change
	router.HandlerFunc(http.MethodGet, "/v1/entries/:id", app.showEntryHandler) //change

	return router
}