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
	pr, err := p.db.CreatePR(ctx, prID, prName, authorID)
	return pr, errs.ParsePgErr(err)
}

func (p *prs) Merge() {

}

func (p *prs) Reassign() {

}
