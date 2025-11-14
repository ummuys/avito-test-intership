package repository

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/config"
	"github.com/ummuys/avito-test-intership/internal/models"
)

type prDB struct {
	logger *zerolog.Logger
	pool   *pgxpool.Pool
}

func NewPRDB(ctx context.Context, logger *zerolog.Logger) (PRDB, error) {
	dbctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cfg, err := config.ParsePRDBEnv()
	if err != nil {
		return nil, err
	}

	pool, err := PoolFromConfig(dbctx, cfg, "pr")
	if err != nil {
		return nil, err
	}

	return &prDB{pool: pool, logger: logger}, nil
}

func (p *prDB) GetReview() {}

func (p *prDB) CreatePR(ctx context.Context, prID, prName, authorID string) (resp models.PRResponse, err error) {
	p.logger.Debug().Str("evt", "call CreatePR").Msg("")
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	var tx pgx.Tx
	tx, err = p.pool.Begin(dbCtx)
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(ctx); rbErr != nil && !errors.Is(rbErr, pgx.ErrTxClosed) {
				p.logger.Error().Err(rbErr).Msg("rollback failed")
			}
		}
	}()

	_, err = tx.Exec(dbCtx, CreatePRStep1, prID, prName, authorID)
	if err != nil {
		saveRawErr(p.logger, "CreatePRStep1", err)
		return
	}

	rews := make([]string, 0, 2)

	var rows pgx.Rows
	rows, err = tx.Query(dbCtx, CreatePRStep2, authorID)
	if err != nil {
		saveRawErr(p.logger, "CreatePRStep2", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r string
		err = rows.Scan(&r)
		if err != nil {
			return
		}
		rews = append(rews, r)
	}

	if err = rows.Err(); err != nil {
		return
	}

	if err = tx.Commit(ctx); err != nil {
		return
	}

	resp.PRID = prID
	resp.AuthorID = authorID
	resp.PRName = prName
	resp.Status = "OPEN"
	resp.AssignedReviewers = append(resp.AssignedReviewers, rews...)

	return resp, nil

}
func (p *prDB) MarkPRAsMerge() {}
func (p *prDB) ReassignPR()    {}
