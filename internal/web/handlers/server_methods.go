package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/service"
)

type sh struct {
	svc    service.ServerService
	logger *zerolog.Logger
}

func NewServerHandler(svc service.ServerService, logger *zerolog.Logger) ServerHandler {
	return &sh{svc: svc, logger: logger}
}

func (s *sh) Health(g *gin.Context) {

}
