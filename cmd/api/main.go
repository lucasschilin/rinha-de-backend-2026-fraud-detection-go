package main

import (
	"log"
	"net/http"

	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/config"
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/router"
)

func main() {
	cfg := config.Load()

	r := router.New()

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	log.Printf("API listening on :%s", cfg.Port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
