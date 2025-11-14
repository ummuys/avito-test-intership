package service

import (
	"context"

	"github.com/ummuys/avito-test-intership/internal/models"
)

type TeamService interface {
	Add(ctx context.Context, body models.AddTeamRequest) error
	Get(ctx context.Context, teamName string) (models.GetTeamResponse, error)
}
