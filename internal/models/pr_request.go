package models

type CreatePRRequest struct {
	PRID     string `json:"pull_request_id"`
	PRName   string `json:"pull_request_name"`
	AuthorID string `json:"author_id"`
}

type MergePRRequest struct {
	PRID string `json:"pull_request_id"`
}
