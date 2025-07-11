package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/Hanumath1006/skillsync/handlers"
)

func main() {
    http.HandleFunc("/healthz", handlers.HealthCheck)
    http.HandleFunc("/register", handlers.Register)

    fmt.Println("Server running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

