package search

import (
	"math"

	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/dataset"
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/vector"
)

const K = 5

type Neighbor struct {
	Dist  float32
	Label uint8
}

func FindKNN(query vector.Vector, ds *dataset.Dataset) [K]Neighbor {
	var closestNeighbors [K]Neighbor
	for cn := range closestNeighbors {
		closestNeighbors[cn].Dist = math.MaxFloat32
	}

	for v := 0; v < len(ds.Vectors); v++ {
		dist := distanceSquared(query, ds.Vectors[v])

		furtherIndex := -1
		furtherDist := float32(-1)

		for j := 0; j < K; j++ {
			if closestNeighbors[j].Dist > furtherDist {
				furtherDist = closestNeighbors[j].Dist
				furtherIndex = j
			}
		}

		if dist < furtherDist {
			closestNeighbors[furtherIndex] = Neighbor{
				Dist:  dist,
				Label: ds.Labels[v],
			}
		}
	}

	return closestNeighbors
}

func Score(neighbors [K]Neighbor) float64 {
	var fraud int

	for n := 0; n < K; n++ {
		if neighbors[n].Label == 1 {
			fraud++
		}
	}

	return float64(fraud) / K
}

func distanceSquared(a, b vector.Vector) float32 {
	var sum float32

	for d := 0; d < len(a); d++ {
		diff := a[d] - b[d]
		sum += diff * diff
	}

	return sum
}
