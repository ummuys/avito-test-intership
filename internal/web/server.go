package web

import (
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/di"
	"github.com/ummuys/avito-test-intership/internal/web/middleware"
)

func InitServer(hand di.Handlers, sec di.Secure, logger *zerolog.Logger) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()

	prh := hand.PRHandler
	th := hand.TeamHandler
	uh := hand.UserHandler
	sh := hand.ServerHandler
	auh := hand.AuthHandler
	adh := hand.AdminHandler

	tm := sec.TokenManager

	allUsers := []string{"admin", "user"}
	onlyAdmin := []string{"admin"}

	// ------- GROUPS  -------

	// MAIN
	api := g.Group("")
	api.Use(middleware.RequestLogger(logger))
	api.Use(gin.Recovery())

	// ALL USERS
	publicGroup := api.Group("")

	// ONLY AUTH USERS
	authGroup := api.Group("")
	authGroup.Use(middleware.Auth(tm, allUsers))

	// ONLY ADMINS
	adminGroup := api.Group("")
	adminGroup.Use(middleware.Auth(tm, onlyAdmin))

	// ------- PUBLIC GROUP -------
	// TEAMS
	publicGroup.POST(createTeamPath, th.Create)

	// SERVER
	publicGroup.GET(healthPath, sh.Health)

	// AUTH
	publicGroup.POST(authPath, auh.Authorization)
	publicGroup.GET(updateAccessToken, auh.UpdateAccessToken)

	// ------- AUTH GROUP  -------
	// TEAMS
	authGroup.GET(getTeamPath, th.Get)

	// USERS
	authGroup.GET(getUserReviewPath, uh.GetReviews)

	// ------- ADMIN GROUP -------
	// USERS
	adminGroup.POST(setUserActivePath, uh.SetState)

	// PR
	adminGroup.POST(createPRPath, prh.Create)
	adminGroup.POST(mergePRPath, prh.Merge)
	adminGroup.POST(reassignPRPath, prh.Reassign)

	// ADMIN
	adminGroup.POST(createSvcUserPath, adh.CreateUser)

	host := os.Getenv("SERVER_IP")
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:              net.JoinHostPort(host, port),
		Handler:           g,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	return server
}

func RunServer(server *http.Server) error {
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
