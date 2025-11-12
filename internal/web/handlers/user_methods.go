package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/service"
)

type uh struct {
	svc    service.UserService
	logger *zerolog.Logger
}

func NewUserHandler(svc service.UserService, logger *zerolog.Logger) UserHandler {
	return &uh{svc: svc, logger: logger}
}

func (u *uh) SetState(g *gin.Context) {

}

func (u *uh) Get(g *gin.Context) {

}
