package service

import (
	"log"

	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/domain"
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/search"
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/vector"
)

type FraudService struct {
	builder *vector.Builder
	tree    *search.Node
}

func NewFraudService(
	builder *vector.Builder,
	tree *search.Node,
) *FraudService {
	return &FraudService{
		builder: builder,
		tree:    tree,
	}
}

func (s *FraudService) Score(request domain.FraudScoreRequest) domain.FraudScoreResponse {
	v := s.builder.Build(request)
	log.Printf("vectorized payload=%v", v)

	neighbors := search.KNN(s.tree, v, search.K)
	log.Printf("neighbors=%v", neighbors)

	score := search.Score(neighbors)
	log.Printf("score=%v", score)

	return domain.FraudScoreResponse{
		Approved:   score < 0.6,
		FraudScore: score,
	}
}
