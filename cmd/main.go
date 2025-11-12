package main

import (
	"context"
	"log"
	"os/signal"
	"sync"
	"syscall"

	"github.com/ummuys/avito-test-intership/internal/di"
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
	rep := di.InitRepositories(tools.Logger.DbLog)
	svc := di.InitServices(rep, tools.Logger.SvcLog)
	hand := di.InitHandlers(svc, tools.Logger.SrvLog)
	srv := web.InitServer(hand)
	appLogger.Info().Msg("Init all interfaces")

	// SYNC TOOLS
	wg := sync.WaitGroup{}
	errch := make(chan error, 1)

	wg.Go(func() {
		errch <- web.RunServer(srv)
	})

	<-ctx.Done()
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
		tools.Logger.AppLog.Error().Msg("fatal shutdown")
	} else {
		tools.Logger.AppLog.Info().Msg("shutdown successful")
	}

}
