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

	// router.HandlerFunc(http.MethodGet, "/v1/appointments/:id", app.requirePermission("appointments:read", app.listAppointment))
  router.HandlerFunc(http.MethodPost, "/v1/appointments", app.requirePermission("appointments:write", app.createAppointments))
	router.HandlerFunc(http.MethodGet, "/v1/appointments/:id", app.requirePermission("appointments:read", app.showAppointment))
	router.HandlerFunc(http.MethodPatch, "/v1/appointments/:id", app.requirePermission("appointments:read", app.updateAppointment))
	router.HandlerFunc(http.MethodDelete, "/v1/appointments/:id", app.requirePermission("appointments:read", app.deleteAppointment))

	router.HandlerFunc(http.MethodPost, "/v1/user", app.createUser)
	router.HandlerFunc(http.MethodPut, "/v1/user/activated", app.activateUser)

	router.HandlerFunc(http.MethodPost, "/v1/token/authentication", app.createAuthenticationToken)

	// router.HandlerFunc(http.MethodGet, "/v1/user/:id", app.showUser)
	// router.HandlerFunc(http.MethodGet, "/v1/user", app.listUsers)
	// router.HandlerFunc(http.MethodPatch, "/v1/user/", app.updateUser)
	// router.HandlerFunc(http.MethodDelete, "/v1/user/:id", app.deleteUser)

	return app.recoverPanic(app.enableCORS(app.rateLimit(app.authenticate(router))))
}
