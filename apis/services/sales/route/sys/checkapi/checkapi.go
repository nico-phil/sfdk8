package checkapi

import (
	"context"
	"math/rand"
	"net/http"

	"github.com/nico-phil/service/app/api/errs"
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

func testerr(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		return errs.Newf(errs.FailedPrecondition, "this is message is trusted")
	}

	status := struct {
		Status string
	}{
		Status: "OK REDINESS",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}

func testpanic(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		panic("WE ARE PANICKING")
	}

	status := struct {
		Status string
	}{
		Status: "OK REDINESS",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
