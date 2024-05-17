package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET index")
}

type Config struct {
	Port int
}

func main() {
	cfg := &Config{
		Port: 8080,
	}

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", Index)

	fmt.Printf("Serving on port %d\n", cfg.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router)
}
