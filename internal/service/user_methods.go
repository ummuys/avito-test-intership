package service

import (
	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/repository"
)

type us struct {
	db     repository.PRDB
	logger *zerolog.Logger
}

func NewUserService(db repository.PRDB, logger *zerolog.Logger) UserService {
	return &us{db: db, logger: logger}
}

func (u *us) SetState() {

}

func (u *us) Get() {

}
