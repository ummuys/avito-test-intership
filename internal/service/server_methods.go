package service

import (
	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/repository"
)

type ss struct {
	logger *zerolog.Logger
}

func NewServerService(db repository.PRDB, logger *zerolog.Logger) ServerService {
	return &ss{logger: logger}
}

func (s *ss) Health() {

}
