package models

type Project struct {
    ID             int      `json:"id"`
    Title          string   `json:"title"`
    Description    string   `json:"description"`
    RequiredSkills []string `json:"required_skills"`
    OwnerID        int      `json:"owner_id"`
}

var Projects []Project
