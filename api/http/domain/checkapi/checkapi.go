package checkapi

import (
	"context"
	"net/http"
	"time"

	"github.com/nico-phil/service/fondation/logger"
	"github.com/nico-phil/service/fondation/web"
)

type api struct {
	build string
	log   *logger.Logger
}

func NewAPI(build string, log *logger.Logger) *api {
	return &api{
		build: build,
		log:   log,
	}
}

func (api *api) Readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	status := "ok"
	statusCode := http.StatusOK

	data := struct {
		Status string `json:"status"`
	}{
		Status: status,
	}

	return web.Respond(ctx, w, data, statusCode)

}
