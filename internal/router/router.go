package router

import (
	"net/http"

	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/handler"
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/service"
)

func New() http.Handler {
	mux := http.NewServeMux()

	fraudService := service.NewFraudService()
	fraudHandler := handler.NewFraudHandler(fraudService)

	mux.HandleFunc("GET /ready", handler.Ready)
	mux.HandleFunc("POST /fraud-score", fraudHandler.Score)

	return mux
}
