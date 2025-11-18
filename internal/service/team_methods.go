package service

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/errs"
	"github.com/ummuys/avito-test-intership/internal/models"
	"github.com/ummuys/avito-test-intership/internal/repository"
)

type ts struct {
	db     repository.TeamDB
	logger *zerolog.Logger
}

func NewTeamService(db repository.TeamDB, logger *zerolog.Logger) TeamService {
	return &ts{db: db, logger: logger}
}

func (t *ts) Add(ctx context.Context, body models.AddTeamRequest) error {
	t.logger.Debug().Str("evt", "call AddTeam").Msg("")
	return errs.ParsePgErr(t.db.AddTeam(ctx, body))
}

func (t *ts) Get(ctx context.Context, teamName string) (models.GetTeamResponse, error) {
	t.logger.Debug().Str("evt", "call GetTeam").Msg("")
	team, err := t.db.GetTeam(ctx, teamName)
	return team, errs.ParsePgErr(err)
}
