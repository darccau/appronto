package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/darccau/appronto/internal/data"
)

func (app *application) createUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Password  string `json:"password"`
		Email     string `json:"email"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIdParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
	}

	user := data.User{
		Id:        id,
		FirstName: "Eduardo",
		LastName:  "Paixao",
		Email:     "darccau@gmail.com",
		CreatedAt: time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
	if err != nil {
		app.logger.Print(err)
		app.serverErrorResponse(w, r, err)
	}
}
