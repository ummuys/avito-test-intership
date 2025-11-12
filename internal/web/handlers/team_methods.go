package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/service"
)

type th struct {
	svc    service.TeamService
	logger *zerolog.Logger
}

func NewTeamHandler(svc service.TeamService, logger *zerolog.Logger) TeamHandler {
	return &th{svc: svc, logger: logger}
}

func (t *th) Create(g *gin.Context) {

}

func (t *th) Get(g *gin.Context) {

}
