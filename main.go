package main

import (
	"go-color-pages/handlers"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Use PORT from environment (Render) or default to 8080 for local
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := chi.NewRouter()

	// Built-in logging middleware
	r.Use(middleware.Logger)

	// Serve static files like CSS
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Routes
	r.Get("/", handlers.PageHandler)              // Page 1
	r.Get("/page/{number}", handlers.PageHandler) // Pages 2-4

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
