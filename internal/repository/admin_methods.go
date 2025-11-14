package repository

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/config"
)

type adDB struct {
	logger *zerolog.Logger
	pool   *pgxpool.Pool
}

func NewAdminDB(ctx context.Context, logger *zerolog.Logger) (AdminDB, error) {
	dbctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cfg, err := config.ParseAdminDBEnv()
	if err != nil {
		return nil, err
	}

	pool, err := PoolFromConfig(dbctx, cfg, "admin")
	if err != nil {
		return nil, err
	}

	return &adDB{pool: pool, logger: logger}, nil
}

func (ad *adDB) CreateUser(ctx context.Context, username string, hashPassword string, role string) (err error) {
	ad.logger.Debug().Str("evt", "call CreateUser").Msg("")
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	var tx pgx.Tx
	tx, err = ad.pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(dbCtx); rbErr != nil && !errors.Is(rbErr, pgx.ErrTxClosed) {
				ad.logger.Error().Err(rbErr).Msg("rollback failed")
			}
		}
	}()

	_, err = tx.Exec(dbCtx, NewUserStep1, username, hashPassword)
	if err != nil {
		saveRawErr(ad.logger, "NewUserStep1", err)
		return
	}

	_, err = tx.Exec(dbCtx, NewUserStep2, username, role)
	if err != nil {
		saveRawErr(ad.logger, "NewUserStep2", err)
		return
	}

	if err = tx.Commit(dbCtx); err != nil {
		return
	}

	return
}

func (ad *adDB) ValidateRole(ctx context.Context, role string) error {
	ad.logger.Debug().Str("evt", "call CheckRole").Msg("")
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	_, err := ad.pool.Exec(dbCtx, CheckRole, role)
	if err != nil {
		saveRawErr(ad.logger, "CheckRole", err)
	}
	return err
}
