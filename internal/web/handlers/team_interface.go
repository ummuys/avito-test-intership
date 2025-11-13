package handlers

import "github.com/gin-gonic/gin"

type TeamHandler interface {
	Add(g *gin.Context)
	Get(g *gin.Context)
}
