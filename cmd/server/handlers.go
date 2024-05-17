package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (app *application) getItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", app.models.ItemsModel.GetItems())
}

func (app *application) getItemID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "itemID")
	intId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintf(w, "%s is not a valid integer ID", id)
	}

	fmt.Fprintf(w, "%v", app.models.ItemsModel.GetItem(intId))
}
