package service

import "github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/domain"

type FraudService struct{}

func NewFraudService() *FraudService {
	return &FraudService{}
}

func (s *FraudService) Score(_ domain.FraudScoreRequest) domain.FraudScoreResponse {
	return domain.FraudScoreResponse{
		Approved:   true,
		FraudScore: 0.0,
	}
}
