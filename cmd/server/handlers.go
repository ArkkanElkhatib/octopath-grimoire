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
		Sort:        "-name",
		Page:        1,
		PageSize:    25,
	}

	items := app.models.ItemModel.GetItems(filter)
	data := envelope{"items": items, "filter": filter}

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
	data := envelope{
		"item": item,
	}

	err = app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.writeJSON(w, http.StatusNotFound, envelope{"error": "item not found"}, nil)
	}
}

func (app *application) getEquipments(w http.ResponseWriter, r *http.Request) {
	var filter data.EquipmentFilter
	filter = data.EquipmentFilter{
		Query:       "",
		QueryTarget: "",
		Extension:   "",
		Sort:        "name",
		Page:        1,
		PageSize:    25,
	}

	equipments := app.models.EquipmentModel.GetEquipments(filter)
	data := envelope{"equipments": equipments, "filter": filter}

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
	data := envelope{
		"equipment": equipment,
	}

	err = app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.writeJSON(w, http.StatusNotFound, envelope{"error": "equipment not found"}, nil)
	}
}
