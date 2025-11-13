package service

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/errs"
	"github.com/ummuys/avito-test-intership/internal/repository"
	"github.com/ummuys/avito-test-intership/internal/secure"
)

type admService struct {
	logger *zerolog.Logger
	db     repository.AdminDB
	ph     secure.PasswordHasher
}

func NewAdminService(logger *zerolog.Logger, db repository.AdminDB, ph secure.PasswordHasher) AdminService {
	return &admService{logger: logger, db: db, ph: ph}
}

func (a *admService) CreateUser(pCtx context.Context, username, password, role string) error {
	a.logger.Debug().Str("evt", "call CreateUser").Msg("")

	err := a.db.ValidateRole(pCtx, role)
	if err != nil {
		return errs.ParsePgErr(err)
	}

	hashPass, err := a.ph.Hash(password)
	if err != nil {
		return err
	}

	if err := a.db.CreateUser(pCtx, username, hashPass, role); err != nil {
		return errs.ParsePgErr(err)
	}

	return nil
}

func (a *admService) UpdateUser(pCtx context.Context, userID int64, username, password, role string) error {
	a.logger.Debug().Str("evt", "call CreateUser").Msg("")

	var (
		err      error
		hashPass string
	)

	if password != "" {
		hashPass, err = a.ph.Hash(password)
		if err != nil {
			return err
		}
	}

	if role != "" {
		err = a.db.ValidateRole(pCtx, role)
		if err != nil {
			return errs.ParsePgErr(err)
		}
	}

	if err := a.db.UpdateUser(pCtx, userID, username, hashPass, role); err != nil {
		return errs.ParsePgErr(err)
	}

	return nil
}

func (a *admService) DeleteUser(pCtx context.Context, username string) error {
	a.logger.Debug().Str("evt", "call DeleteUser").Msg("")

	if err := a.db.DeleteUser(pCtx, username); err != nil {
		return errs.ParsePgErr(err)
	}

	return nil
}

// TO FIX: MAKE BETTER ERR CHECKER
func (a *admService) GetUser(pCtx context.Context, userID int64) (, error) {
	a.logger.Debug().Str("evt", "call GetUsers").Msg("")

	list, err := a.db.GetUsers(pCtx)
	if err != nil {
		return nil, errs.ParsePgErr(err)
	}
	return list, nil
}
