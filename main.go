package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Hanumath1006/skillsync/handlers"
	"github.com/Hanumath1006/skillsync/middleware"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/healthz", handlers.HealthCheck)
	router.HandleFunc("/register", handlers.Register)
	router.HandleFunc("/login", handlers.Login)

	router.Handle("/me", middleware.AuthMiddleware(http.HandlerFunc(handlers.Me)))
	router.Handle("/users", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetUsers)))

	router.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetProjects(w, r)
		} else if r.Method == http.MethodPost {
			middleware.AuthMiddleware(http.HandlerFunc(handlers.CreateProject)).ServeHTTP(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	router.Handle("/match/{projectId}", middleware.AuthMiddleware(http.HandlerFunc(handlers.MatchUsers)))

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
