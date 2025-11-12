package service

import (
	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/repository"
)

type ts struct {
	db     repository.PRDB
	logger *zerolog.Logger
}

func NewTeamService(db repository.PRDB, logger *zerolog.Logger) TeamService {
	return &ts{db: db, logger: logger}
}

func (t *ts) Create() {

}

func (t *ts) Get() {

}
