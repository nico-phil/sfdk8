package checkapi

import (
	"context"
	"net/http"

	"github.com/nico-phil/service/fondation/web"
)

func liveness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "OK LIVENESS",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}

func readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "OK REDINESS",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
