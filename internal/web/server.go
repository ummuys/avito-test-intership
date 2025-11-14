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

	//allUsers := []string{"admin", "user"}
	onlyAdmin := []string{"admin"}

	// MAIN
	api := g.Group("")
	api.Use(middleware.RequestLogger(logger))
	api.Use(gin.Recovery())

	// TEAMS
	teams := api.Group("")
	teams.POST(createTeamPath, th.Create)
	teams.GET(getTeamPath, th.Get)

	// USERS
	admUsers := api.Group("")
	admUsers.Use(middleware.Auth(tm, onlyAdmin))
	admUsers.POST(setUserActivePath, uh.SetState)

	users := api.Group("")
	users.GET(getUserReviewPath, uh.Get)

	// AUTH
	auth := api.Group("")
	auth.POST(authPath, auh.Authorization)
	auth.GET(updateAccessToken, auh.UpdateAccessToken)

	// ADMIN
	adm := api.Group("")
	adm.Use(middleware.Auth(tm, onlyAdmin))
	adm.POST(createSvcUserPath, adh.CreateUser)

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
