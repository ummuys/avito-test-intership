package service

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/errs"
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
	username, team_name, err := u.db.SetUserState(ctx, userID, state)
	return username, team_name, errs.ParsePgErr(err)
}

func (u *us) Get() {

}
