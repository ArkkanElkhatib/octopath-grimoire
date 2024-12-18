package main

import (
	_ "net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) setupRoutes() {
	app.setupMiddleware()
	app.setupItemRoutes("/items")
	app.setupEquipmentRoutes("/equipments")
}

func (app *application) setupMiddleware() {
	app.router.Use(middleware.Logger)
}

func (app *application) setupItemRoutes(prefix string) {
	app.router.Route(prefix, func(r chi.Router) {
		r.Get("/", app.getItems)
		r.Get("/{itemID}", app.getItemID)
	})
}

func (app *application) setupEquipmentRoutes(prefix string) {
	app.router.Route(prefix, func(r chi.Router) {
		r.Get("/", app.getEquipments)
		r.Get("/{equipmentID}", app.getEquipmentId)
	})
}
