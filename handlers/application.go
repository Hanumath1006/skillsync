package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Hanumath1006/skillsync/middleware"
	"github.com/Hanumath1006/skillsync/models"
	"github.com/gorilla/mux"
)

// POST /apply/{projectId}
func ApplyToProject(w http.ResponseWriter, r *http.Request) {
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

	// Find the project
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

	// Save applicant
	models.Applications[projectID] = append(models.Applications[projectID], *user)

	fmt.Fprintf(w, "✅ Applied to project %d successfully", projectID)
}

// GET /applicants/{projectId}
func ViewApplicants(w http.ResponseWriter, r *http.Request) {
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

	// Check if the user owns the project
	var ownsProject bool
	for _, p := range models.Projects {
		if p.ID == projectID && p.OwnerID == user.UserID {
			ownsProject = true
			break
		}
	}

	if !ownsProject {
		http.Error(w, "Forbidden: You do not own this project", http.StatusForbidden)
		return
	}

	applicants := models.Applications[projectID]
	json.NewEncoder(w).Encode(applicants)
}

