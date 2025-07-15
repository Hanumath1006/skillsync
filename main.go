package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Hanumath1006/skillsync/handlers"
	"github.com/Hanumath1006/skillsync/middleware"
)

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/healthz", handlers.HealthCheck)
    mux.HandleFunc("/register", handlers.Register)
    mux.HandleFunc("/login", handlers.Login)

    // ✅ Protect /me route with JWT middleware
    mux.Handle("/me", middleware.AuthMiddleware(http.HandlerFunc(handlers.Me)))

    fmt.Println("Server running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", mux))
}