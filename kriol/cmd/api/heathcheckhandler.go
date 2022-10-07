// Filename: cmd/api/healthcheck.go

package main

import (

	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	//Create a map to hold our healthcheck data
	data := map[string]string {
		"status": "available",
		"environment": app.config.env,
		"version": version,
	}
	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server made a whoopsie and couldn't process, bro", http.StatusInternalServerError)
		return
	}
}