package repository

import (
	"context"

	"github.com/ummuys/avito-test-intership/internal/models"
)

type PRDB interface {
	// PR
	GetReview()
	Create(ctx context.Context, prID, prName, authorID string) (resp models.PRResponse, err error)
	Merge(ctx context.Context, prID string) (models.MergeRPResponse, error)
	ReassignPR()
}
