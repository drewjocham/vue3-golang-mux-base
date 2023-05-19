package main

import (
	"github.com/interviews/internal/data"
	validator "github.com/interviews/internal/vaildator"
	"net/http"
)

type mockData struct {
	Name     string `json:"firstName"`
	Lastname string `json:"lastName"`
}

func (app *application) listTestHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title  string
		Genres []string
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Title = app.readString(qs, "title", "")
	input.Genres = app.readCSV(qs, "genres", []string{})

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafeList = []string{"id", "title", "year", "runtime", "-id", "-title", "-year", "-runtime"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	tempData := []*mockData{
		{Name: "Drew", Lastname: "Jocham"},
		{Name: "Peter", Lastname: "vue"},
		{Name: "GoLang", Lastname: "lang"},
		{Name: "AI", Lastname: "rules"},
	}

	err := app.writeJSON(w, http.StatusOK, envelope{"data": tempData, "metadata": "none"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
