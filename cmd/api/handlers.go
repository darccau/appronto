package main

import (
	"fmt"
	"net/http"
)

func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "What the fuck is that shit")
}

func (app *application) showUser(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintln(w, id)
}
