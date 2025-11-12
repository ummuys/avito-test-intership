package di

import (
	"github.com/ummuys/avito-test-intership/internal/config"
	"github.com/ummuys/avito-test-intership/internal/repository"
	"github.com/ummuys/avito-test-intership/internal/service"
	"github.com/ummuys/avito-test-intership/internal/web/handlers"
)

type Tools struct {
	Logger *config.Loggers
}

type Repositories struct {
	PRDB repository.PRDB
}

type Services struct {
	PRService     service.PRService
	ServerService service.ServerService
	TeamService   service.TeamService
	UserService   service.UserService
}

type Handlers struct {
	PRHandler     handlers.PRHandler
	ServerHandler handlers.ServerHandler
	TeamHandler   handlers.TeamHandler
	UserHandler   handlers.UserHandler
}
