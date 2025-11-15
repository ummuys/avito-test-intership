package models

type SetUserStateRequest struct {
	UserID   string `json:"user_id"`
	IsActive bool   `json:"is_active"`
}

type GetUserReviewsRequest struct {
	UserID string `json:"user_id"`
}
