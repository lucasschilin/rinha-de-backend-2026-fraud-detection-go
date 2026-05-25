package search

import "github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/dataset"

func Build(ds *dataset.MmapDataset) ([]Node, int) {
	nodes := make([]Node, ds.Count)

	var build func(start, end int) int

	build = func(start, end int) int {
		if start >= end {
			return -1
		}

		root := start

		nodes[root].Index = root

		if end-start == 1 {
			nodes[root].Left = -1
			nodes[root].Right = -1
			return root
		}

		mid := (start + end) / 2

		rootVector, _, _ := ds.GetRecord(root)
		midVector, _, _ := ds.GetRecord(mid)

		nodes[root].Radius = Distance(rootVector, midVector)

		nodes[root].Left = build(start+1, mid)
		nodes[root].Right = build(mid, end)

		return root
	}

	root := build(0, ds.Count)

	return nodes, root
}
