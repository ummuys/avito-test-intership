package service

import (
	"context"

	"github.com/ummuys/avito-test-intership/internal/models"
)

type PRService interface {
	Create(ctx context.Context, prID, prName, authorID string) (models.PRResponse, error)
	Merge()
	Reassign()
}
