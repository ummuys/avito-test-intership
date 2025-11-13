package web

import (
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/di"
	"github.com/ummuys/avito-test-intership/internal/web/middleware"
)

func InitServer(hand di.Handlers, logger *zerolog.Logger) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()

	prh := hand.PRHandler
	th := hand.TeamHandler
	uh := hand.UserHandler
	sh := hand.ServerHandler

	// MAIN
	api := g.Group("")
	api.Use(middleware.RequestLogger(logger))
	api.Use(gin.Recovery())

	// TEAMS
	teams := api.Group("")
	teams.POST(createTeamPath, th.Add)
	teams.GET(getTeamPath, th.Get)

	// USERS
	users := api.Group("")
	users.POST(setUserActivePath, uh.SetState)
	users.GET(getUserReviewPath, uh.Get)

	// PR
	pr := api.Group("")
	pr.POST(createPRPath, prh.Create)
	pr.POST(mergePRPath, prh.Merge)
	pr.POST(reassignPRPath, prh.Reassign)

	//SERVER
	srv := api.Group("")
	srv.GET(healthPath, sh.Health)

	host := os.Getenv("SERVER_IP")
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: g,
	}

	return server
}

func RunServer(server *http.Server) error {
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
