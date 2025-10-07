package mux

import (
	"context"
	"fmt"

	"net/http"

	"github.com/nico-phil/service/fondation/logger"
	"github.com/nico-phil/service/fondation/web"
)

func WebAPI(log *logger.Logger) *web.App {

	logger := func(ctx context.Context, msg string, v ...any) {
		log.Info(ctx, msg, v...)
	}
	app := web.NewApp(logger)

	f := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
		fmt.Fprintln(w, "hello world")
		return nil
	}
	app.HandleFunc("/hello", f)

	return app
}
