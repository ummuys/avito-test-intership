package repository

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/config"
	"github.com/ummuys/avito-test-intership/internal/models"
)

type tDB struct {
	logger *zerolog.Logger
	pool   *pgxpool.Pool
}

func NewTeamDB(ctx context.Context, logger *zerolog.Logger) (TeamDB, error) {
	dbctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cfg, err := config.ParseTeamDBEnv()
	if err != nil {
		return nil, err
	}

	pool, err := PoolFromConfig(dbctx, cfg, "report")
	if err != nil {
		return nil, err
	}

	return &tDB{pool: pool, logger: logger}, nil
}

// TEAM
func (t *tDB) AddTeam(ctx context.Context, body models.AddTeamRequest) (err error) {
	t.logger.Debug().Str("evt", "call AddTeam").Msg("")

	dbCtx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	var tx pgx.Tx
	tx, err = t.pool.Begin(dbCtx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(ctx); rbErr != nil && !errors.Is(rbErr, pgx.ErrTxClosed) {
				t.logger.Error().Err(rbErr).Msg("rollback failed")
			}
		}
	}()

	b := &pgx.Batch{}
	teamID := uuid.New()
	b.Queue(AddTeamQuery, teamID, body.TeamName)
	for _, m := range body.Members {
		b.Queue(AddUserQuery, m.UserID, m.Username, teamID, m.IsActive)
	}

	sb := tx.SendBatch(dbCtx, b)
	for i := 0; i < len(body.Members)+1; i++ {
		if _, err = sb.Exec(); err != nil {
			_ = sb.Close()
			return err
		}
	}

	if err = sb.Close(); err != nil {
		return
	}

	if err = tx.Commit(ctx); err != nil {
		return
	}

	return nil
}

func (t *tDB) GetTeam(ctx context.Context, teamName string) ([][]any, error) {
	t.logger.Debug().Str("evt", "call AddTeam").Msg("")

	dbCtx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	rows, err := t.pool.Query(dbCtx, GetTeamQuery, teamName)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var mbrs [][]any
	for rows.Next() {
		vals, err := rows.Values()
		if err != nil {
			return nil, err
		}
		row := make([]any, len(vals))
		copy(row, vals)
		mbrs = append(mbrs, row)
	}

	if mbrs == nil {
		return nil, pgx.ErrNoRows
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return mbrs, nil
}
