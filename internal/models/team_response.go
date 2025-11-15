package models

type GetTeamResponse struct {
	TeamName string   `json:"team_name"`
	Members  []Member `json:"members"`
}
