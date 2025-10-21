package models

type User struct {
	Id       string `json:"id,omitempty"`
	Role     string `json:"role,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type RegisterUser struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"Str0ngP@ss!"`
}
