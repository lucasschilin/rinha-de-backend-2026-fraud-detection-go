package handler

import (
	"encoding/json"
	"net/http"

	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/domain"
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/service"
)

type FraudHandler struct {
	service *service.FraudService
}

func NewFraudHandler(service *service.FraudService) *FraudHandler {
	return &FraudHandler{
		service: service,
	}
}

func (h *FraudHandler) Score(w http.ResponseWriter, r *http.Request) {
	var request domain.FraudScoreRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	response := h.service.Score(request)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(response)
}
