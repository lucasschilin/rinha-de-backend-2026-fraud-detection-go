package service

import (
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/dataset"
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/domain"
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/search"
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/vector"
)

type FraudService struct {
	builder *vector.Builder
	nodes   []search.Node
	dataset *dataset.MmapDataset
	root    int
}

func NewFraudService(
	builder *vector.Builder,
	nodes []search.Node,
	dataset *dataset.MmapDataset,
	root int,
) *FraudService {
	return &FraudService{
		builder: builder,
		nodes:   nodes,
		dataset: dataset,
		root:    root,
	}
}

func (s *FraudService) Score(request domain.FraudScoreRequest) domain.FraudScoreResponse {
	v := s.builder.Build(request)

	neighbors := search.KNN(s.nodes, s.dataset, s.root, v, search.K)

	score := search.Score(neighbors)

	return domain.FraudScoreResponse{
		Approved:   score < 0.6,
		FraudScore: score,
	}
}
