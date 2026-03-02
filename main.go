package main

import (
	"go-color-pages/handlers"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// recoverMiddleware from panics in downstream handlers,
// logs the error, and returns a 500 Internal Server Error
func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("PANIC RECOVERED: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Use PORT from environment (Render) or default to 8080 for local
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := chi.NewRouter()

	// Built-in logging middleware
	r.Use(recoverMiddleware)
	r.Use(middleware.Logger)

	// Serve static files like CSS
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Routes
	r.Get("/", handlers.PageHandler)              // Page 1
	r.Get("/page/{number}", handlers.PageHandler) // Pages 2-4

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
