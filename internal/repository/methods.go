package repository

import (
	"github.com/rs/zerolog"
)

type prDB struct {
	logger *zerolog.Logger
}

func NewPRDB(logger *zerolog.Logger) PRDB {
	return &prDB{logger: logger}
}
