package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ismail-hafiane/goinfra/internal/database"
	"github.com/ismail-hafiane/goinfra/internal/handlers"
	"github.com/ismail-hafiane/goinfra/internal/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	database.Connect()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.HealthCheck)
	mux.HandleFunc("/api/users", handlers.Users)
	mux.Handle("/metrics", promhttp.Handler())

	handler := middleware.Logging(mux)
	handler = middleware.Recovery(handler)

	fmt.Printf("🚀 Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
