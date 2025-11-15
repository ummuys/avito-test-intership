package models

type UserWrapper struct {
	User User `json:"user"`
}

type User struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	TeamName string `json:"team_name"`
	IsActive bool   `json:"is_active"`
}

type GetUserReviewsResponse struct {
	UserID string   `json:"user_id"`
	PR     []UserPR `json:"pull_requests"`
}

type UserPR struct {
	PRID     string `json:"pull_request_id"`
	PRName   string `json:"pull_request_name"`
	AuthorID string `json:"author_id"`
	Status   string `json:"status"`
}
