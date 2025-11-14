package main

import (
	"context"
	"errors"
	"log"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/ummuys/avito-test-intership/internal/config"
	"github.com/ummuys/avito-test-intership/internal/di"
	"github.com/ummuys/avito-test-intership/internal/errs"
	"github.com/ummuys/avito-test-intership/internal/web"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	tools, err := di.InitTools()
	if err != nil {
		log.Fatal(err)
	}

	appLogger := tools.Logger.AppLog

	// INTERFACES
	rep, err := di.InitRepositories(ctx, tools.Logger.DbLog)
	if err != nil {
		appLogger.Fatal().Err(err).Msg("")
	}
	sec, err := di.InitSecure()
	if err != nil {
		appLogger.Fatal().Err(err).Msg("")
	}
	svc := di.InitServices(rep, sec, tools.Logger.SvcLog)
	hand := di.InitHandlers(svc, sec, tools.Logger.SrvLog)
	srv := web.InitServer(hand, sec, tools.Logger.SrvLog)
	appLogger.Info().Msg("Init all interfaces")

	appConf, err := config.ParseAppConfig()
	if err != nil {
		appLogger.Fatal().Err(err).Msg("load app config failed")
	}

	if err := svc.AdminService.CreateUser(ctx, appConf.Username, appConf.Password, "admin"); err == nil || errors.Is(err, errs.ErrPGDuplicate) {
		tools.Logger.DbLog.Info().Msg("default admin user initialized")
	} else {
		tools.Logger.DbLog.Error().Err(err).Msg("failed to init admin user")
	}

	// SYNC TOOLS
	wg := sync.WaitGroup{}
	errch := make(chan error, 2)

	wg.Go(func() {
		<-ctx.Done()
		defer srv.Close()
		srvCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		errch <- srv.Shutdown(srvCtx)
	})

	wg.Go(func() {
		errch <- web.RunServer(srv)
	})

	wg.Wait()
	close(errch)

	haveErr := false
	for err := range errch {
		if err != nil {
			appLogger.Error().Err(err).Msg("")
			haveErr = true
		}
	}

	if haveErr {
		tools.Logger.AppLog.Error().Msg("Fatal shutdown")
	} else {
		tools.Logger.AppLog.Info().Msg("Shutdown successful")
	}

}
