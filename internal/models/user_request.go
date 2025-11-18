package models

type SetUserStateRequest struct {
	UserID   string `json:"user_id" binding:"required"`
	IsActive bool   `json:"is_active" binding:"required"`
}

type GetUserReviewsRequest struct {
	UserID string `json:"user_id" binding:"required"`
}
