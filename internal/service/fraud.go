package service

import (
	"log"

	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/domain"
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/search"
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/vector"
)

type FraudService struct {
	builder *vector.Builder
	nodes   []search.Node
	records []search.Record
	root    int
}

func NewFraudService(
	builder *vector.Builder,
	nodes []search.Node,
	records []search.Record,
	root int,
) *FraudService {
	return &FraudService{
		builder: builder,
		nodes:   nodes,
		records: records,
		root:    root,
	}
}

func (s *FraudService) Score(request domain.FraudScoreRequest) domain.FraudScoreResponse {
	v := s.builder.Build(request)

	log.Printf("vectorized payload=%v", v)

	neighbors := search.KNN(s.nodes, s.records, s.root, v, search.K)

	log.Printf("neighbors=%v", neighbors)

	score := search.Score(neighbors)

	return domain.FraudScoreResponse{
		Approved:   score < 0.6,
		FraudScore: score,
	}
}
