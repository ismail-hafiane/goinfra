package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ismail-hafiane/goinfra/internal/handlers"
	"github.com/ismail-hafiane/goinfra/internal/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/health", handlers.HealthCheck)
	mux.HandleFunc("/api/users", handlers.Users)

	// Middleware
	handler := middleware.Logging(mux)
	handler = middleware.Recovery(handler)

	fmt.Printf("🚀 Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
