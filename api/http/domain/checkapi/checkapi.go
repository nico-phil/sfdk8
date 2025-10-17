package checkapi

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/nico-phil/service/app/api/errs"
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

func (api *api) readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	status := "ok"
	statusCode := http.StatusOK

	data := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  status,
		Message: "readiness",
	}

	return web.Respond(ctx, w, data, statusCode)
}

func (api *api) liveness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	status := "ok"
	statusCode := http.StatusOK

	data := struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}{
		Status:  status,
		Message: "liveness",
	}

	return web.Respond(ctx, w, data, statusCode)

}

func (api *api) testerr(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// ctx, cancel := context.WithTimeout(ctx, time.Second)
	// defer cancel()

	// status := "unknown"
	// statusCode := http.StatusBadRequest

	// data := struct {
	// 	Status  string `json:"status"`
	// 	Message string `json:"message"`
	// }{
	// 	Status:  status,
	// 	Message: "test error",
	// }

	return errs.New(errs.FailedPrecondition, errors.New("this is a test error to make sure it works"))

	// web.Respond(ctx, w, data, statusCode)

}
