package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"fmt"

	"github.com/Hanumath1006/skillsync/middleware"
	"github.com/Hanumath1006/skillsync/models"
	"github.com/gorilla/mux"
)

func MatchUsers(w http.ResponseWriter, r *http.Request) {
	user := middleware.GetUserFromContext(r)
	if user == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	projectID, err := strconv.Atoi(vars["projectId"])
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	var project *models.Project
	for i := range models.Projects {
		if models.Projects[i].ID == projectID {
			project = &models.Projects[i]
			break
		}
	}

	if project == nil {
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

	fmt.Println("✅ Project found:", project.Title)
    fmt.Println("🔍 Matching against skills:", project.RequiredSkills)

	var matchedUsers []models.User
	for _, u := range models.Users {
		fmt.Println("👤 Checking user:", u.Email, "with skills:", u.Skills)
		if hasMatchingSkills(u.Skills, project.RequiredSkills) {
			fmt.Println("🎯 Matched user:", u.Email)
			matchedUsers = append(matchedUsers, u)
		}
	}

	if matchedUsers == nil {
    	matchedUsers = []models.User{} // avoid null in JSON
    }

	json.NewEncoder(w).Encode(matchedUsers)
}

func hasMatchingSkills(userSkills, requiredSkills []string) bool {
	userSkillSet := make(map[string]bool)
	for _, s := range userSkills {
		clean := strings.ToLower(strings.TrimSpace(s))
		userSkillSet[clean] = true
	}

	for _, rs := range requiredSkills {
		cleaned := strings.ToLower(strings.TrimSpace(rs))
		if userSkillSet[cleaned] {
			fmt.Println("✅ Skill matched:", cleaned)
			return true
		}
	}
	return false
}
