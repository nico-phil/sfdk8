package main

import (
	"context"
	"os"
	"runtime"

	"github.com/nico-phil/service/fondation/logger"
)

func main() {
	var log *logger.Logger

	events := logger.Events{
		Error: func(ctx context.Context, r logger.Record) {
			log.Error(ctx, "****** SEND ALETER *******")
		},
	}

	traceIDFn := func(ctx context.Context) string {
		return "223"
	}

	log = logger.NewWithEvents(os.Stdout, logger.LevelInfo, "SALES", traceIDFn, events)

	ctx := context.Background()
	err := run(ctx, log)
	if err != nil {
		log.Error(ctx, "startup", "msg", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, log *logger.Logger) error {
	log.Info(ctx, "startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))
	return nil
}
