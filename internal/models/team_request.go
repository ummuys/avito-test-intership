package models

type AddTeamRequest struct {
	TeamName string   `json:"team_name"`
	Members  []Member `json:"members"`
}

type Member struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}
