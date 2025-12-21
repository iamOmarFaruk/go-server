package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"go-server/internal/handlers"
)

func main() {
	// Initialize page handler
	pageHandler, err := handlers.NewPageHandler("templates")
	if err != nil {
		log.Fatal("Failed to create page handler:", err)
	}

	// Create router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.RealIP)
	r.Use(middleware.Compress(5))

	// Static files
	staticFs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", staticFs))

	// Routes
	r.Get("/", pageHandler.Home)
	r.Get("/about", pageHandler.About)
	r.Get("/testimonials", pageHandler.Testimonials)
	
	// 404 handler - must be last
	r.NotFound(pageHandler.NotFound)

	// Health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}