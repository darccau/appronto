package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/user", app.createUser)
	router.HandlerFunc(http.MethodGet, "/v1/user/:id", app.showUser)
	router.HandlerFunc(http.MethodGet, "/v1/user", app.listUsers)
	router.HandlerFunc(http.MethodPatch, "/v1/user/:id", app.updateUser)
	router.HandlerFunc(http.MethodDelete, "/v1/user/:id", app.deleteUser)

	return app.rateLimit(app.recoverPanic(router))
}
