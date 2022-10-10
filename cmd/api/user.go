package main

import (
	"net/http"
	"time"

	"github.com/darccau/appronto/internal/data"
)

func (app *application) showUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)
	if err != nil {
		http.NotFound(w, r)
	}

	user := data.User{
		Id:        id,
		FirstName: "Eduardo",
		LastName:  "Paixao",
		Email:     "darccau@gmail.com",
		CreatedAt: time.Now(),
	}

	err = app.writeJson(w, http.StatusOK, envelope{"user": user}, nil)
	if err != nil {
		app.logger.Print(err)
		http.Error(w, "The server encoutered a problem and could not process your request", http.StatusInternalServerError)
	}
}
