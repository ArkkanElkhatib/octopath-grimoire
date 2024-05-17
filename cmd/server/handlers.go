package main

import (
	"fmt"
	"net/http"
	"strconv"

	"octopath-grimoire/internal/data"

	"github.com/go-chi/chi/v5"
)

func (app *application) getItems(w http.ResponseWriter, r *http.Request) {
	var filter data.ItemFilter
	filter = data.ItemFilter{
		Query:       "",
		QueryTarget: "",
		Extension:   "",
		Sort:        "",
		Page:        3,
		PageSize:    200,
	}

	fmt.Fprintf(w, "%v + %d", app.models.ItemsModel.GetItems(filter), len(app.models.ItemsModel.GetItems(filter)))
}

func (app *application) getItemID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "itemID")
	intId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintf(w, "%s is not a valid integer ID", id)
	}

	fmt.Fprintf(w, "%v", app.models.ItemsModel.GetItem(intId))
}
