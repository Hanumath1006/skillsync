package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/Hanumath1006/skillsync/models"
	"github.com/Hanumath1006/skillsync/utils"
)

var users []models.User
var userIDCounter = 1

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input models.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	input.ID = userIDCounter
	userIDCounter++
	input.Password = hashedPassword

	users = append(users, input)

	// Don't return the hashed password
	input.Password = ""
	json.NewEncoder(w).Encode(input)

	fmt.Printf("Registered user: %+v\n", input)
}
