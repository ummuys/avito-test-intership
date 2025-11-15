package handlers

import "github.com/gin-gonic/gin"

type AdminHandler interface {
	CreateUser(g *gin.Context)
}
