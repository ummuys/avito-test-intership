package repository

import (
	"context"

	"github.com/ummuys/avito-test-intership/internal/models"
)

type PRDB interface {
	// PR
	Create(ctx context.Context, prID, prName, authorID string) (resp models.PRResponse, err error)
	Merge(ctx context.Context, prID string) (models.MergeRPResponse, error)
	ReassignPR(ctx context.Context, prID string, oldUserID string) (resp models.ReassignPRResponse, err error)
}
