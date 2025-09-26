package checkapi

import (
	"encoding/json"
	"net/http"
)

func liveness(w http.ResponseWriter, r *http.Request) {
	status := struct {
		Status string
	}{
		Status: "OK LIVENESS",
	}

	json.NewEncoder(w).Encode(status)
}

func readiness(w http.ResponseWriter, r *http.Request) {
	status := struct {
		Status string
	}{
		Status: "OK REDINESS",
	}

	json.NewEncoder(w).Encode(status)
}
