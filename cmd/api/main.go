package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const port = ":9999"

func readyHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]bool{
		"ok": true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(response)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ready", readyHandler)

	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	log.Println("API listening on " + port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
