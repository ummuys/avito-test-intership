package handlers

import "github.com/gin-gonic/gin"

type AuthHandler interface {
	UpdateAccessToken(g *gin.Context)
	Authorization(g *gin.Context)
}
