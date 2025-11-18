package models

type AddTeamRequest struct {
	TeamName string   `json:"team_name" binding:"required"`
	Members  []Member `json:"members" binding:"required"`
}

type Member struct {
	UserID   string `json:"user_id" binding:"required"`
	Username string `json:"username" binding:"required"`
	IsActive bool   `json:"is_active" binding:"required"`
}
