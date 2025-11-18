package service

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/errs"
	"github.com/ummuys/avito-test-intership/internal/models"
	"github.com/ummuys/avito-test-intership/internal/repository"
)

type prs struct {
	db     repository.PRDB
	logger *zerolog.Logger
}

func NewPRService(db repository.PRDB, logger *zerolog.Logger) PRService {
	return &prs{db: db, logger: logger}
}

func (p *prs) Create(ctx context.Context, prID, prName, authorID string) (models.PRResponse, error) {
	p.logger.Debug().Str("evt", "call CreatePR").Msg("")
	pr, err := p.db.Create(ctx, prID, prName, authorID)
	return pr, errs.ParsePgErr(err)
}

func (p *prs) Merge(ctx context.Context, prID string) (models.MergeRPResponse, error) {
	p.logger.Debug().Str("evt", "call MergePR").Msg("")
	pr, err := p.db.Merge(ctx, prID)
	return pr, errs.ParsePgErr(err)
}

func (p *prs) Reassign(ctx context.Context, prID, oldUserID string) (models.ReassignPRResponse, error) {
	p.logger.Debug().Str("evt", "call ReassignPR").Msg("")
	pr, err := p.db.ReassignPR(ctx, prID, oldUserID)
	return pr, errs.ParsePgErr(err)
}
