package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/user", app.createUser)
	router.HandlerFunc(http.MethodGet, "/v1/user/:id", app.showUser)
	router.HandlerFunc(http.MethodPut, "/v1/user/:id", app.updateUser)
	router.HandlerFunc(http.MethodDelete, "/v1/user/:id", app.deleteUser)

	return router
}
