package handler

import (
	"encoding/json"
	"net/http"
)

func Ready(w http.ResponseWriter, _ *http.Request) {
	response := map[string]bool{
		"ok": true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(response)
}
