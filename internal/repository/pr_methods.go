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

func (p *prDB) Create(ctx context.Context, prID, prName, authorID string) (resp models.PRResponse, err error) {
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

	rows.Close()

	b := &pgx.Batch{}
	for i, r := range rews {
		b.Queue(CreatePRStep3, prID, r, i+1)
	}

	sb := tx.SendBatch(dbCtx, b)
	for i := 0; i < len(rews); i++ {
		if _, err = sb.Exec(); err != nil {
			saveRawErr(p.logger, "CreatePRStep3", err)
			_ = sb.Close()
			return
		}
	}
	sb.Close()

	if err = tx.Commit(ctx); err != nil {
		return
	}

	resp.PRID = prID
	resp.AuthorID = authorID
	resp.PRName = prName
	resp.Status = "OPEN"
	if len(rews) != 0 {
		resp.AssignedReviewers = append(resp.AssignedReviewers, rews...)
	}

	return resp, nil

}
func (p *prDB) Merge(ctx context.Context, prID string) (resp models.MergeRPResponse, err error) {
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
			if rbErr := tx.Rollback(dbCtx); rbErr != nil && !errors.Is(rbErr, pgx.ErrTxClosed) {
				p.logger.Error().Err(rbErr).Msg("rollback failed")
			}
		}
	}()

	var status string
	err = tx.QueryRow(dbCtx, CheckIsPRMerge, prID).Scan(&status)
	if status != "MERGED" {

		_, err = tx.Exec(dbCtx, MergePRStep1, prID)
		if err != nil {
			saveRawErr(p.logger, "CreatePRStep1", err)
			return models.MergeRPResponse{}, err
		}

	}

	var (
		prName   string
		authorID string
		mergedAt time.Time
	)

	err = tx.QueryRow(dbCtx, MergePRStep2, prID).Scan(&prName, &authorID, &mergedAt)
	if err != nil {
		return
	}

	var rows pgx.Rows
	rows, err = tx.Query(dbCtx, MergePRStep3, prID)
	if err != nil {
		return
	}
	defer rows.Close()

	rews := make([]string, 0, 2)
	for rows.Next() {
		var r string
		rows.Scan(&r)
		rews = append(rews, r)
	}

	if rows.Err() != nil {
		return
	}

	if err = tx.Commit(dbCtx); err != nil {
		return
	}

	resp.PRID = prID
	resp.PRName = prName
	resp.AuthorID = authorID
	resp.Status = "MERGED"
	resp.MergeAt = mergedAt
	if len(rews) != 0 {
		resp.AssignedReviewers = append(resp.AssignedReviewers, rews...)
	}

	return resp, nil
}
func (p *prDB) ReassignPR() {}
