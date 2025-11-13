package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/config"
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

	pool, err := PoolFromConfig(dbctx, cfg, "report")
	if err != nil {
		return nil, err
	}

	return &prDB{pool: pool, logger: logger}, nil
}

func (p *prDB) GetReview()     {}
func (p *prDB) CreatePR()      {}
func (p *prDB) MarkPRAsMerge() {}
func (p *prDB) ReassignPR()    {}
