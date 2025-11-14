package repository

import (
	"context"

	"github.com/ummuys/avito-test-intership/internal/models"
)

type PRDB interface {
	// PR
	GetReview()
	CreatePR(ctx context.Context, prID, prName, authorID string) (resp models.PRResponse, err error)
	MarkPRAsMerge()
	ReassignPR()
}
