package router

import (
	"net/http"

	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/dataset"
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/handler"
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/risk"
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/service"
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/vector"
)

func New(ds *dataset.Dataset) http.Handler {
	mux := http.NewServeMux()

	fraudService := service.NewFraudService(
		vector.NewBuilder(risk.NewDefault()), ds,
	)
	fraudHandler := handler.NewFraudHandler(fraudService)

	mux.HandleFunc("GET /ready", handler.Ready)
	mux.HandleFunc("POST /fraud-score", fraudHandler.Score)

	return mux
}
