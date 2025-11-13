package repository

type PRDB interface {
	// PR
	GetReview()
	CreatePR()
	MarkPRAsMerge()
	ReassignPR()
}
