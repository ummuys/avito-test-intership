package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type sh struct {
	logger *zerolog.Logger
}

func NewServerHandler(logger *zerolog.Logger) ServerHandler {
	return &sh{logger: logger}
}

func (s *sh) Health(g *gin.Context) {
	s.logger.Debug().Str("evt", "call Health").Msg("")
	g.Set("msg", "server ok")
	g.JSON(http.StatusOK, gin.H{})
}
