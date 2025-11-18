package service

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/errs"
	"github.com/ummuys/avito-test-intership/internal/models"
	"github.com/ummuys/avito-test-intership/internal/repository"
)

type us struct {
	db     repository.UserDB
	logger *zerolog.Logger
}

func NewUserService(db repository.UserDB, logger *zerolog.Logger) UserService {
	return &us{db: db, logger: logger}
}

func (u *us) SetState(ctx context.Context, userID string, state bool) (string, string, error) {
	u.logger.Debug().Str("evt", "call SetState").Msg("")
	username, team_name, err := u.db.SetUserState(ctx, userID, state)
	return username, team_name, errs.ParsePgErr(err)
}

func (u *us) GetReviews(ctx context.Context, userID string) ([]models.UserPR, error) {
	u.logger.Debug().Str("evt", "call GetReviews").Msg("")
	upr, err := u.db.GetReviews(ctx, userID)
	return upr, errs.ParsePgErr(err)
}
