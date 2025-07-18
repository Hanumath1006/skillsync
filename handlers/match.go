package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

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

	var matchedUsers []models.User
	for _, u := range models.Users {
		if hasMatchingSkills(u.Skills, project.RequiredSkills) {
			matchedUsers = append(matchedUsers, u)
		}
	}

	json.NewEncoder(w).Encode(matchedUsers)
}

func hasMatchingSkills(userSkills, requiredSkills []string) bool {
	userSkillSet := make(map[string]bool)
	for _, s := range userSkills {
		userSkillSet[strings.ToLower(s)] = true
	}

	for _, rs := range requiredSkills {
		if userSkillSet[strings.ToLower(rs)] {
			return true
		}
	}
	return false
}
