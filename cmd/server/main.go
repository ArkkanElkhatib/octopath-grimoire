package main

import (
	"fmt"
	"net/http"

	"octopath-grimoire/internal/data"

	"github.com/go-chi/chi/v5"
)

type config struct {
	Port          int
	ItemsFilepath string
}

type application struct {
	router *chi.Mux
	models *data.Models
	config *config
}

func main() {
	appCfg := &config{
		Port: 8080,
	}

	modelsCfg := &data.ModelsConfig{
		ItemsFilepath: "./assets/data/octopath_items.csv",
	}

	app := &application{
		router: chi.NewRouter(),
		models: data.LoadModels(modelsCfg),
	}

	app.setupRoutes()

	fmt.Printf("Serving on port %d\n", appCfg.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", appCfg.Port), app.router)
}
