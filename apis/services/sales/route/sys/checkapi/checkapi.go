package checkapi

import (
	"context"
	"encoding/json"
	"net/http"
)

func liveness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "OK LIVENESS",
	}

	return json.NewEncoder(w).Encode(status)
}

func readiness(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "OK REDINESS",
	}

	return json.NewEncoder(w).Encode(status)
}
