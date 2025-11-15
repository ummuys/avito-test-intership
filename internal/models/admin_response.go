package models

type CreateUserResponse struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}
