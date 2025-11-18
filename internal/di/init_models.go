package di

import (
	"github.com/ummuys/avito-test-intership/internal/config"
	"github.com/ummuys/avito-test-intership/internal/repository"
	"github.com/ummuys/avito-test-intership/internal/secure"
	"github.com/ummuys/avito-test-intership/internal/service"
	"github.com/ummuys/avito-test-intership/internal/web/handlers"
)

type Tools struct {
	Logger *config.Loggers
}

type Repositories struct {
	PRDB    repository.PRDB
	UserDB  repository.UserDB
	TeamDB  repository.TeamDB
	AdminDB repository.AdminDB
	AuthDB  repository.AuthDB
}

type Secure struct {
	PasswordHasher secure.PasswordHasher
	TokenManager   secure.TokenManager
}

type Services struct {
	PRService    service.PRService
	TeamService  service.TeamService
	UserService  service.UserService
	AdminService service.AdminService
	AuthService  service.AuthService
}

type Handlers struct {
	PRHandler     handlers.PRHandler
	ServerHandler handlers.ServerHandler
	TeamHandler   handlers.TeamHandler
	UserHandler   handlers.UserHandler
	AdminHandler  handlers.AdminHandler
	AuthHandler   handlers.AuthHandler
}
