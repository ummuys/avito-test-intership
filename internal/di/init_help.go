package di

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/logger"
	"github.com/ummuys/avito-test-intership/internal/repository"
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

func InitRepositories(logger *zerolog.Logger) Repositories {
	prdb := repository.NewPRDB(logger)
	return Repositories{PRDB: prdb}
}

func InitServices(rep Repositories, logger *zerolog.Logger) Services {
	prs := service.NewPRService(rep.PRDB, logger)
	ss := service.NewServerService(rep.PRDB, logger)
	ts := service.NewTeamService(rep.PRDB, logger)
	us := service.NewUserService(rep.PRDB, logger)
	return Services{PRService: prs, ServerService: ss, TeamService: ts, UserService: us}
}

func InitHandlers(svc Services, logger *zerolog.Logger) Handlers {
	prh := handlers.NewPRHandler(svc.PRService, logger)
	sh := handlers.NewServerHandler(svc.ServerService, logger)
	th := handlers.NewTeamHandler(svc.TeamService, logger)
	uh := handlers.NewUserHandler(svc.UserService, logger)
	return Handlers{PRHandler: prh, ServerHandler: sh, TeamHandler: th, UserHandler: uh}
}
