package errs

import "errors"

var (
	ErrNotEnoughReviewers = errors.New("not enough reviewers assigned")
	ErrPRMerged           = errors.New("PR_MERGED")
	ErrRVNotAssigned      = errors.New("NOT_ASSIGNED")
	ErrNoCandidate        = errors.New("NO_CANDIDATE")
)

const (
	ErrCodeNotEnoughReviewers = "NOT_ENOUGH_REVIEWERS"
	ErrMsgNotEnoughReviewers  = "There must be at least one reviewer assigned"
)

const (
	ErrCodePRMerged = "PR_MERGED"
	ErrMsgPRMerged  = "cannot reassign on merged PR"
)

const (
	ErrCodeRVNotAssigned = "NOT_ASSIGNED"
	ErrMsgRVNotAssigned  = "reviewer is not assigned to this PR"
)

const (
	ErrCodeNoCandidate = "NO_CANDIDATE"
	ErrMsgNoCandidate  = "no active replacement candidate in team"
)

const (
	ErrCodePRExists = "PR_EXISTS"
	ErrMsgPRExists  = "PR id already exists"
)
