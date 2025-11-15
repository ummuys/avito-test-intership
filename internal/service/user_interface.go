package service

import (
	"context"

	"github.com/ummuys/avito-test-intership/internal/models"
)

type UserService interface {
	SetState(ctx context.Context, userID string, state bool) (string, string, error)
	GetReviews(ctx context.Context, userID string) ([]models.UserPR, error)
}
