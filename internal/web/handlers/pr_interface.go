package handlers

import "github.com/gin-gonic/gin"

type PRHandler interface {
	Create(g *gin.Context)
	Merge(g *gin.Context)
	Reassign(g *gin.Context)
}
