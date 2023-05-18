package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	//router.HandlerFunc(http.MethodGet, "/v1/test", app.requirePermission("test:read", app.listTestHandler))

	router.HandlerFunc(http.MethodGet, "/v1/test", app.listTestHandler)

	return app.metrics(app.recoverPanic(app.enableCORS(app.rateLimit(app.authenticate(router)))))
}
