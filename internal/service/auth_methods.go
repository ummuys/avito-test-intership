package service

import (
	"context"
	"errors"

	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/errs"
	"github.com/ummuys/avito-test-intership/internal/repository"
	"github.com/ummuys/avito-test-intership/internal/secure"
)

type as struct {
	logger *zerolog.Logger
	db     repository.AdminDB
	ph     secure.PasswordHasher
}

func NewAuthService(logger *zerolog.Logger, db repository.AdminDB, ph secure.PasswordHasher) AuthService {
	return &as{logger: logger, db: db, ph: ph}
}

func (a *as) CheckCredentials(ctx context.Context, username, password string) (int64, string, error) {
	a.logger.Debug().Str("evt", "call CheckCredentials").Msg("")

	user_id, role, hashPass, err := a.db.CheckCredentials(ctx, username)
	if err != nil {
		return 0, "", errs.ParsePgErr(err)
	}

	if !u.ph.CheckHash(password, hashPass) {
		return 0, "", errors.New(errs.ErrCodeInvalidTeamName)
	}

	return user_id, role, nil
}
