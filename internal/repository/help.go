package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/config"
)

func PoolFromConfig(ctx context.Context, config config.DBConfig, dbName string) (*pgxpool.Pool, error) {
	poolCfg, err := pgxpool.ParseConfig(config.Addr)
	if err != nil {
		return nil, err
	}
	poolCfg.MinConns = config.MinConn
	poolCfg.MaxConns = config.MaxConn
	poolCfg.MaxConnLifetime = config.MaxConnLifetime
	poolCfg.MaxConnLifetimeJitter = config.MaxConnLifetimeJitter
	poolCfg.MaxConnIdleTime = config.MaxConnIdleTime
	poolCfg.HealthCheckPeriod = config.HealthCheckPeriod

	var conn *pgxpool.Pool
	for i := 0; i < 5; i++ {
		conn, err = pgxpool.NewWithConfig(ctx, poolCfg)
		if err == nil {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}

	if err != nil {
		return nil, fmt.Errorf("can't connect to db: %w", err)
	}

	if err = conn.Ping(ctx); err != nil {
		return nil, fmt.Errorf("%s_db didn't pinged: %w", dbName, err)
	}
	return conn, nil
}

func saveRawErr(logger *zerolog.Logger, queryName string, err error) {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		msg := fmt.Sprintf("%s failed", queryName)
		logger.Error().
			Str("pg_code", pgErr.Code).
			Str("pg_message", pgErr.Message).
			Str("pg_detail", pgErr.Detail).
			Msg(msg)
	} else {
		msg := fmt.Sprintf("%s failed (non pg error)", queryName)
		logger.Error().Err(err).Msg(msg)
	}
}
