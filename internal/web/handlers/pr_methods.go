package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/service"
)

type prh struct {
	svc    service.PRService
	logger *zerolog.Logger
}

func NewPRHandler(svc service.PRService, logger *zerolog.Logger) PRHandler {
	return &prh{svc: svc, logger: logger}
}

func (p *prh) Create(g *gin.Context) {

}

func (p *prh) Merge(g *gin.Context) {

}

func (p *prh) Reassign(g *gin.Context) {

}
