package handlers

import "github.com/gin-gonic/gin"

type ServerHandler interface {
	Health(g *gin.Context)
}
