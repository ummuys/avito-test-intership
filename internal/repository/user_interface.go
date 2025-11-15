package repository

import (
	"context"

	"github.com/ummuys/avito-test-intership/internal/models"
)

type UserDB interface {
	SetUserState(ctx context.Context, userID string, state bool) (string, string, error)
	GetReviews(ctx context.Context, userID string) ([]models.UserPR, error)
}
