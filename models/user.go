package models

type User struct {
    ID       int      `json:"id"`
    Name     string   `json:"name"`
    Email    string   `json:"email"`
    Password string   `json:"-"`
    Skills   []string `json:"skills"`
}

var Users []User
