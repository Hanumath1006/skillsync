package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/Hanumath1006/skillsync/models"
    "github.com/Hanumath1006/skillsync/middleware"
    "strconv"
)

var (
    projects        []models.Project
    projectIDCounter = 1
)

// POST /projects
func CreateProject(w http.ResponseWriter, r *http.Request) {
    user := middleware.GetUserFromContext(r)
    if user == nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    var input struct {
        Title         string   `json:"title"`
        Description   string   `json:"description"`
        RequiredSkills []string `json:"required_skills"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid body", http.StatusBadRequest)
        return
    }

    newProject := models.Project{
        ID:             projectIDCounter,
        Title:          input.Title,
        Description:    input.Description,
        RequiredSkills: input.RequiredSkills,
        OwnerID:        user.UserID,
    }
    projectIDCounter++
    projects = append(projects, newProject)

    json.NewEncoder(w).Encode(newProject)
}

// GET /projects
func GetProjects(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(projects)
}
