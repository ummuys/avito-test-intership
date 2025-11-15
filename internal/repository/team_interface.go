package repository

import (
	"context"

	"github.com/ummuys/avito-test-intership/internal/models"
)

type TeamDB interface {
	AddTeam(ctx context.Context, body models.AddTeamRequest) error
	GetTeam(ctx context.Context, teamName string) (models.GetTeamResponse, error)
}
