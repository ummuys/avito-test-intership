package handlers

import "github.com/gin-gonic/gin"

type TeamHandler interface {
	Create(g *gin.Context)
	Get(g *gin.Context)
}
