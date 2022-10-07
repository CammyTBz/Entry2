// Filename/cmd/api/entry.go

package main

import (
	"fmt"
	"net/http"
	"time"

	"kriol.camerontillett.net/internal/data"
)

//createEntryHandler for the "POST /v1/entry" endpoint
func (app *application) createEntryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new entry..")
}

//createEntryHandler for the "GET /v1/entry/:id" endpoint
func (app *application) showEntryHandler(w http.ResponseWriter, r *http.Request) {
	// Get the value of the "id" parameter
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Create a new instance of the Entries struct containing the ID we extracted
	//from our URL and some sample data
	entry := data.Entry {
		ID: id,
		CreatedAt: time.Now(),
		Name: "Yo Mama",
		Level: "High School",
		Contact: "Inita Lyfe",
		Phone: "666-7777",
		Address: "14 Upyoaph Street",
		Mode: []string{"blended", "online"},
		Version: 1,
	}

	err = app.writeJSON(w, http.StatusOK, entry, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "Ain't work, dawg", http.StatusInternalServerError)
	}
}