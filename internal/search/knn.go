package search

import (
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/dataset"
)

const K = 5

type Neighbor struct {
	Dist  float32
	Label uint8
}

func KNN(
	nodes []Node,
	ds *dataset.MmapDataset,
	root int,
	target [14]float32,
	k int,
) []Neighbor {
	h := NewMaxHeap(k)

	var search func(i int)

	search = func(i int) {
		if i == -1 {
			return
		}

		n := nodes[i]

		vector, label, _ := ds.GetRecord(n.Index)

		d := Distance(target, vector)

		h.Push(Neighbor{
			Label: label,
			Dist:  d,
		})

		if n.Left == -1 && n.Right == -1 {
			return
		}

		var first, second int

		if d < n.Radius {
			first = n.Left
			second = n.Right
		} else {
			first = n.Right
			second = n.Left
		}

		if first != -1 {
			search(first)
		}

		if second != -1 {
			if d-n.Radius <= h.worst() {
				search(second)
			}
		}
	}

	search(root)

	return h.Items()
}

func Score(neighbors []Neighbor) float64 {
	var fraudWeight float64
	var totalWeight float64

	for _, n := range neighbors {
		w := 1.0 / (float64(n.Dist) + 0.0001)

		totalWeight += w

		if n.Label == 1 {
			fraudWeight += w
		}
	}

	return fraudWeight / totalWeight
}
