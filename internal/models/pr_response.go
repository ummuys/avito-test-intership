package models

import "time"

type CreatePRWrapper struct {
	PR PRResponse `json:"pr"`
}

type PRResponse struct {
	PRID              string   `json:"pull_request_id"`
	PRName            string   `json:"pull_request_name"`
	AuthorID          string   `json:"author_id"`
	Status            string   `json:"status"`
	AssignedReviewers []string `json:"assigned_reviewers"`
}

type MergeRPResponse struct {
	PRResponse
	MergeAt time.Time `json:"mergedAt"`
}
type MergeRPWrapper struct {
	PR MergeRPResponse `json:"pr"`
}
