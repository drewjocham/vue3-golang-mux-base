package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/v1/healthcheck", app.healthcheckHandler).Methods("GET")
	// courses
	router.HandleFunc("/v1/courses", app.courses.CoursesAllHandler).Methods("GET")
	router.HandleFunc("/v1/course/{id}", app.courses.CoursesIdHandler).Methods("GET")
	router.HandleFunc("/v1/course", app.courses.CreateCourseHandler).Methods("POST")

	//course details

	return app.middleware.Metrics(app.middleware.RecoverPanic(
		app.middleware.EnableCORS(app.middleware.RateLimit(app.middleware.Authenticate(router)))))
}
