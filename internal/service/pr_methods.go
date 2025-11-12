package service

import (
	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/repository"
)

type prs struct {
	db     repository.PRDB
	logger *zerolog.Logger
}

func NewPRService(db repository.PRDB, logger *zerolog.Logger) PRService {
	return &prs{db: db, logger: logger}
}

func (p *prs) Create() {

}

func (p *prs) Merge() {

}

func (p *prs) Reassign() {

}
