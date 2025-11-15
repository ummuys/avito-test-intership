package service

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/errs"
	"github.com/ummuys/avito-test-intership/internal/repository"
	"github.com/ummuys/avito-test-intership/internal/secure"
)

type ads struct {
	logger *zerolog.Logger
	db     repository.AdminDB
	ph     secure.PasswordHasher
}

func NewAdminService(db repository.AdminDB, ph secure.PasswordHasher, logger *zerolog.Logger) AdminService {
	return &ads{logger: logger, db: db, ph: ph}
}

func (ad *ads) CreateUser(pCtx context.Context, username, password, role string) error {
	ad.logger.Debug().Str("evt", "call CreateUser").Msg("")

	err := ad.db.ValidateRole(pCtx, role)
	if err != nil {
		return errs.ParsePgErr(err)
	}

	hashPass, err := ad.ph.Hash(password)
	if err != nil {
		return err
	}

	if err := ad.db.CreateUser(pCtx, username, hashPass, role); err != nil {
		return errs.ParsePgErr(err)
	}

	return nil
}
