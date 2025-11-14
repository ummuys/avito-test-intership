package models

// {
//   "pr": {
//     "pull_request_id": "pr-1001",
//     "pull_request_name": "Add search",
//     "author_id": "u1",
//     "status": "OPEN",
//     "assigned_reviewers": [
//       "u2",
//       "u3"
//     ]
//   }
// }

type CreatePRResponse struct {
	PR PRResponse `json:"pr"`
}

type PRResponse struct {
	PRID              string   `json:"pull_request_id"`
	PRName            string   `json:"pull_request_name"`
	AuthorID          string   `json:"author_id"`
	Status            string   `json:"status"`
	AssignedReviewers []string `json:"assigned_reviewers"`
}
