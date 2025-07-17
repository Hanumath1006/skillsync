package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Hanumath1006/skillsync/models"
	"github.com/Hanumath1006/skillsync/utils"
	"github.com/Hanumath1006/skillsync/middleware"
)

var users []models.User
var userIDCounter = 1

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Name     string   `json:"name"`
		Email    string   `json:"email"`
		Password string   `json:"password"`
		Skills   []string `json:"skills"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// Construct a new user with the hashed password
	newUser := models.User{
		ID:       userIDCounter,
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
		Skills:   input.Skills,
	}

	userIDCounter++
	users = append(users, newUser)

	// Return response without password
	newUser.Password = ""
	json.NewEncoder(w).Encode(newUser)

	fmt.Printf("Registered user: %+v\n", newUser)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Look for user in memory
	var foundUser *models.User
	for _, user := range users {
		if user.Email == input.Email {
			foundUser = &user
			break
		}
	}

	if foundUser == nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Check password
	if !utils.CheckPasswordHash(input.Password, foundUser.Password) {
		http.Error(w, "Incorrect password", http.StatusUnauthorized)
		return
	}

	// Generate JWT
	token, err := utils.GenerateJWT(foundUser.ID, foundUser.Email)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Respond with token
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func Me(w http.ResponseWriter, r *http.Request) {
    user := middleware.GetUserFromContext(r)
    if user == nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    json.NewEncoder(w).Encode(map[string]any{
        "user_id": user.UserID,
        "email":   user.Email,
    })
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
    user := middleware.GetUserFromContext(r)
    if user == nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    var publicUsers []map[string]interface{}
    for _, u := range users {
        publicUsers = append(publicUsers, map[string]interface{}{
            "id":     u.ID,
            "name":   u.Name,
            "email":  u.Email,
            "skills": u.Skills,
        })
    }

    json.NewEncoder(w).Encode(publicUsers)
}

