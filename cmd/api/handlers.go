package main

import (
	"fmt"
	"net/http"
	"strconv"

	"octopath-grimoire/internal/data"

	"github.com/go-chi/chi/v5"
)

func (app *application) getItems(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()

	querySearch := app.readQueryString(queryValues, "query", "")
	var queryTarget string
	if querySearch != "" {
		queryTarget = app.readQueryString(queryValues, "target", "name")
	} else {
		queryTarget = app.readQueryString(queryValues, "target", "")
	}
	queryExtension := app.readQueryString(queryValues, "extension", "")

	querySort := app.readQueryString(queryValues, "sort", "")
	queryPage := app.readQueryStringAsInt(queryValues, "page", 1)
	queryPageSize := app.readQueryStringAsInt(queryValues, "page_size", 25)
	filter := data.Filter{
		Query:       querySearch,
		QueryTarget: queryTarget,
		Extension:   queryExtension,
		Sort:        querySort,
		Page:        queryPage,
		PageSize:    queryPageSize,
	}

	items, numResults := app.models.ItemModel.GetItems(filter)
	metadata := data.Metadata{
		Total:    numResults,
		Returned: len(items),
	}
	data := envelope{
		"metadata": metadata,
		"items":    items,
		"filter":   filter,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.writeJSON(w, http.StatusNotFound, envelope{"error": "could not write items"}, nil)
	}
}

func (app *application) getItemID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "itemID")
	intId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintf(w, "%s is not a valid integer ID", id)
	}

	item := app.models.ItemModel.GetItem(intId)

	if item.IsEmpty() {
		app.writeJSON(w, http.StatusNotFound, envelope{"error": "item not found"}, nil)
		return
	}

	data := envelope{
		"item": item,
	}

	err = app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.writeJSON(w, http.StatusNotFound, envelope{"error": "item not found"}, nil)
	}
}

func (app *application) getEquipments(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()

	querySearch := app.readQueryString(queryValues, "query", "")
	var queryTarget string
	if querySearch != "" {
		queryTarget = app.readQueryString(queryValues, "target", "name")
	} else {
		queryTarget = app.readQueryString(queryValues, "target", "")
	}
	queryExtension := app.readQueryString(queryValues, "extension", "")

	querySort := app.readQueryString(queryValues, "sort", "")
	queryPage := app.readQueryStringAsInt(queryValues, "page", 1)
	queryPageSize := app.readQueryStringAsInt(queryValues, "page_size", 25)
	filter := data.Filter{
		Query:       querySearch,
		QueryTarget: queryTarget,
		Extension:   queryExtension,
		Sort:        querySort,
		Page:        queryPage,
		PageSize:    queryPageSize,
	}

	equipments, numResults := app.models.EquipmentModel.GetEquipments(filter)
	metadata := data.Metadata{
		Total:    numResults,
		Returned: len(equipments),
	}
	data := envelope{
		"metadata":   metadata,
		"equipments": equipments,
		"filter":     filter,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {

		app.writeJSON(w, http.StatusNotFound, envelope{"error": "could not write equipments"}, nil)
	}
}

func (app *application) getEquipmentId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "equipmentID")
	intId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintf(w, "%s is not a valid integer ID", id)
	}

	equipment := app.models.EquipmentModel.GetEquipment(intId)

	if equipment.IsEmpty() {
		app.writeJSON(w, http.StatusNotFound, envelope{"error": "equipment not found"}, nil)
		return
	}

	data := envelope{
		"equipment": equipment,
	}

	err = app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.writeJSON(w, http.StatusNotFound, envelope{"error": "equipment not found"}, nil)
	}
}
