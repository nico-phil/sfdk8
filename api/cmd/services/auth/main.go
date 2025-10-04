package main

import (
	"context"
	"fmt"
	"runtime"

	"github.com/nico-phil/service/fondation/logger"
)

func main() {
	logger := logger.New()
	err := run(context.Background(), logger)
	if err != nil {
		fmt.Println(err)
	}
}

func run(ctx context.Context, log *logger.Logger) error {
	log.Info(ctx, "startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))
	return nil
}
