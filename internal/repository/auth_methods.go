package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/config"
)

type auDB struct {
	logger *zerolog.Logger
	pool   *pgxpool.Pool
}

func NewAuthDB(ctx context.Context, logger *zerolog.Logger) (AuthDB, error) {
	dbctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cfg, err := config.ParseAuthDBEnv()
	if err != nil {
		return nil, err
	}

	pool, err := PoolFromConfig(dbctx, cfg, "auth")
	if err != nil {
		return nil, err
	}

	return &auDB{pool: pool, logger: logger}, nil
}

func (au *auDB) CheckCredentials(ctx context.Context, username string) (int64, string, string, error) {
	au.logger.Debug().Str("evt", "call CheckCredentials").Msg("")
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	var (
		user_id int64
		role    string
		pass    string
	)
	err := au.pool.QueryRow(ctx, GetCredentials, username).Scan(&user_id, &pass, &role)
	if err != nil {
		saveRawErr(au.logger, "GetCredentials", err)
	}
	return user_id, role, pass, err
}
