package handlers

import "github.com/gin-gonic/gin"

type UserHandler interface {
	SetState(g *gin.Context)
	GetReviews(g *gin.Context)
}
