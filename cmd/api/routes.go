package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/create/users", app.createUser)
	router.HandlerFunc(http.MethodGet, "/v1/show/user/:id", app.showUserHandler)

	return router
}
