package di

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/logger"
	"github.com/ummuys/avito-test-intership/internal/repository"
	"github.com/ummuys/avito-test-intership/internal/secure"
	"github.com/ummuys/avito-test-intership/internal/service"
	"github.com/ummuys/avito-test-intership/internal/web/handlers"
)

func InitTools() (Tools, error) {
	logger, err := logger.InitLogger(os.Getenv("LOG_PATH"))
	if err != nil {
		return Tools{}, err
	}
	return Tools{Logger: logger}, nil
}

func InitRepositories(ctx context.Context, logger *zerolog.Logger) (Repositories, error) {
	prDB, err := repository.NewPRDB(ctx, logger)
	if err != nil {
		return Repositories{}, err
	}

	uDB, err := repository.NewUserDB(ctx, logger)
	if err != nil {
		return Repositories{}, err
	}

	tDB, err := repository.NewTeamDB(ctx, logger)
	if err != nil {
		return Repositories{}, err
	}

	auDB, err := repository.NewAuthDB(ctx, logger)
	if err != nil {
		return Repositories{}, err
	}

	adDB, err := repository.NewAdminDB(ctx, logger)
	if err != nil {
		return Repositories{}, err
	}

	return Repositories{PRDB: prDB, UserDB: uDB, TeamDB: tDB, AdminDB: adDB, AuthDB: auDB}, nil
}

func InitServices(rep Repositories, sec Secure, logger *zerolog.Logger) Services {
	prs := service.NewPRService(rep.PRDB, logger)
	ts := service.NewTeamService(rep.TeamDB, logger)
	us := service.NewUserService(rep.UserDB, logger)
	aus := service.NewAuthService(rep.AuthDB, sec.PasswordHasher, logger)
	ads := service.NewAdminService(rep.AdminDB, sec.PasswordHasher, logger)
	return Services{
		PRService:   prs,
		TeamService: ts, UserService: us,
		AdminService: ads, AuthService: aus,
	}
}

func InitSecure() (Secure, error) {
	tm, err := secure.NewTokenManager()
	if err != nil {
		return Secure{}, err
	}
	ph := secure.NewPasswordHasher()
	return Secure{TokenManager: tm, PasswordHasher: ph}, nil
}

func InitHandlers(svc Services, sec Secure, logger *zerolog.Logger) Handlers {
	prh := handlers.NewPRHandler(svc.PRService, logger)
	sh := handlers.NewServerHandler(logger)
	th := handlers.NewTeamHandler(svc.TeamService, logger)
	uh := handlers.NewUserHandler(svc.UserService, logger)
	auh := handlers.NewAuthHandler(sec.TokenManager, svc.AuthService, logger)
	adh := handlers.NewAdminHandler(svc.AdminService, logger)
	return Handlers{
		PRHandler: prh, ServerHandler: sh,
		TeamHandler: th, UserHandler: uh,
		AdminHandler: adh, AuthHandler: auh,
	}
}
