package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/config"
)

type uDB struct {
	logger *zerolog.Logger
	pool   *pgxpool.Pool
}

func NewUserDB(ctx context.Context, logger *zerolog.Logger) (UserDB, error) {
	dbctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cfg, err := config.ParseUserDBEnv()
	if err != nil {
		return nil, err
	}

	pool, err := PoolFromConfig(dbctx, cfg, "user")
	if err != nil {
		return nil, err
	}

	return &uDB{pool: pool, logger: logger}, nil
}

func (u *uDB) SetUserState(ctx context.Context, userID string, state bool) (string, string, error) {
	u.logger.Debug().Str("evt", "call AddTeam").Msg("")
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	_, err := u.pool.Exec(dbCtx, UpdateUserStateQuery, state, userID)
	if err != nil {
		saveRawErr(u.logger, "UpdateUserStateQuery", err)
		return "", "", err
	}

	var (
		username string
		teamName string
	)

	err = u.pool.QueryRow(dbCtx, GetUserInfoQuery, userID).Scan(&username, &teamName)
	if err != nil {
		saveRawErr(u.logger, "GetUserInfoQuery", err)
		return "", "", err
	}

	return username, teamName, nil

}
