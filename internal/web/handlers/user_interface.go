package handlers

import "github.com/gin-gonic/gin"

type UserHandler interface {
	SetState(g *gin.Context)
	Get(g *gin.Context)
}
