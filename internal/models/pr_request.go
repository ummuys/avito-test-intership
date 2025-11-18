package models

type CreatePRRequest struct {
	PRID     string `json:"pull_request_id" binding:"required"`
	PRName   string `json:"pull_request_name" binding:"required"`
	AuthorID string `json:"author_id" binding:"required"`
}

type MergePRRequest struct {
	PRID string `json:"pull_request_id" binding:"required"`
}

type ReassignPRRequest struct {
	PRID      string `json:"pull_request_id" binding:"required"`
	OldUserID string `json:"old_user_id" binding:"required"`
}
