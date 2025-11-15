package repository

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/config"
	"github.com/ummuys/avito-test-intership/internal/errs"
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

	_, err = tx.Exec(dbCtx, CreatePRQuery, prID, prName, authorID)
	if err != nil {
		saveRawErr(p.logger, "CreatePRQuery", err)
		return
	}

	rews := make([]string, 0, 2)

	var rows pgx.Rows
	rows, err = tx.Query(dbCtx, GetAssignedReviewersQuery, authorID, 2)
	if err != nil {
		saveRawErr(p.logger, "GetAssignedReviewersQuery", err)
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

	rows.Close()

	b := &pgx.Batch{}
	for i, r := range rews {
		b.Queue(PasteARtoHistoryQuery, prID, r, i+1)
	}

	sb := tx.SendBatch(dbCtx, b)
	for i := 0; i < len(rews); i++ {
		if _, err = sb.Exec(); err != nil {
			saveRawErr(p.logger, "PasteARtoHistoryQuery", err)
			_ = sb.Close()
			return
		}
	}

	if err = sb.Close(); err != nil {
		return
	}

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

	var c int
	err = tx.QueryRow(dbCtx, CheckReviewersCountQuery, prID).Scan(&c)
	if err != nil {
		saveRawErr(p.logger, "CheckReviewersCountQuery", err)
		return
	}

	if c < 1 {
		err = errs.ErrNotEnoughReviewers
		return
	}

	var status string
	err = tx.QueryRow(dbCtx, CheckIsPRMergeQuery, prID).Scan(&status)
	if err != nil {
		saveRawErr(p.logger, "CheckIsPRMergeQuery", err)
		return
	}

	if status != "MERGED" {
		_, err = tx.Exec(dbCtx, MarkPRAsMergedQuery, prID)
		if err != nil {
			saveRawErr(p.logger, "MarkPRAsMergedQuery", err)
			return models.MergeRPResponse{}, err
		}
	}

	var (
		prName   string
		authorID string
		mergedAt time.Time
	)

	err = tx.QueryRow(dbCtx, GetMergePRInfoQuery, prID).Scan(&prName, &authorID, &mergedAt)
	if err != nil {
		saveRawErr(p.logger, "GetPRInfoQuery", err)
		return
	}

	var rows pgx.Rows
	rows, err = tx.Query(dbCtx, GetPRARQuery, prID)
	if err != nil {
		saveRawErr(p.logger, "GetPRARQuery", err)
		return
	}
	defer rows.Close()

	var rews []string
	for rows.Next() {
		var r string
		if err = rows.Scan(&r); err != nil {
			return
		}
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

	return
}

func (p *prDB) ReassignPR(ctx context.Context, prID string, oldUserID string) (resp models.ReassignPRResponse, err error) {
	p.logger.Debug().Str("evt", "call ReassignPR").Msg("")
	dbCtx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	var tx pgx.Tx
	tx, err = p.pool.Begin(dbCtx)
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(context.Background()); rbErr != nil && !errors.Is(rbErr, pgx.ErrTxClosed) {
				p.logger.Error().Err(rbErr).Msg("rollback failed")
			}
		}
	}()

	// Check pr_status
	var status string
	if err = tx.QueryRow(dbCtx, CheckIsPRMergeQuery, prID).Scan(&status); err != nil {
		saveRawErr(p.logger, "CheckIsPRMergeQuery", err)
		return
	}

	if status == "MERGED" {
		err = errs.ErrPRMerged
		return
	}

	// Check rv on pr
	var rvInPr bool
	if err = tx.QueryRow(dbCtx, CheckRVInPR, oldUserID, prID).Scan(&rvInPr); err != nil {
		saveRawErr(p.logger, "CheckRVInPR", err)
		return
	}

	if !rvInPr {
		err = errs.ErrRVNotAssigned
		return
	}

	var (
		prName   string
		authorID string
	)

	err = tx.QueryRow(dbCtx, GetReassignRInfoQuery, prID).Scan(&prName, &authorID)
	if err != nil {
		saveRawErr(p.logger, "GetReassignRInfoQueryy", err)
		return
	}

	// Try to find candidate
	var newID string
	if err = tx.QueryRow(dbCtx, GetReassignedReviewersQuery, oldUserID, authorID, prID, 1).Scan(&newID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = errs.ErrNoCandidate
			return
		}
		saveRawErr(p.logger, "GetReassignedReviewersQuery", err)
		return
	}

	_, err = tx.Exec(dbCtx, ChangeHistory, newID, oldUserID, prID)
	if err != nil {
		saveRawErr(p.logger, "ChangeHistory", err)
		return
	}

	var rows pgx.Rows
	rows, err = tx.Query(dbCtx, GetPRARQuery, prID)
	if err != nil {
		saveRawErr(p.logger, "GetPRARQuery", err)
		return
	}
	defer rows.Close()

	var rews []string
	for rows.Next() {
		var r string
		if err = rows.Scan(&r); err != nil {
			return
		}
		rews = append(rews, r)
	}

	if err = rows.Err(); err != nil {
		return
	}

	if err = tx.Commit(context.Background()); err != nil {
		return
	}

	resp.PRID = prID
	resp.PRName = prName
	resp.AuthorID = authorID
	resp.Status = status
	resp.AssignedReviewers = append(resp.AssignedReviewers, rews...)
	resp.ReplacedBy = newID

	return
}
