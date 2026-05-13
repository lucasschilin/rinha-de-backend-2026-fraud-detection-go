package service

import (
	"log"

	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/dataset"
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/domain"
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/search"
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/vector"
)

type FraudService struct {
	builder *vector.Builder
	dataset *dataset.Dataset
}

func NewFraudService(
	builder *vector.Builder,
	dataset *dataset.Dataset,
) *FraudService {
	return &FraudService{
		builder: builder,
		dataset: dataset,
	}
}

func (s *FraudService) Score(request domain.FraudScoreRequest) domain.FraudScoreResponse {
	v := s.builder.Build(request)
	log.Printf("vectorized payload=%v", v)

	neighbors := search.FindKNN(v, s.dataset)
	log.Printf("neighbors=%v", neighbors)

	score := search.Score(neighbors)
	log.Printf("score=%v", score)

	return domain.FraudScoreResponse{
		Approved:   score < 0.6,
		FraudScore: score,
	}
}
