package search

import "sort"

func Build(records []Record) *Node {
	if len(records) == 0 {
		return nil
	}

	node := &Node{
		Point: records[0],
	}

	if len(records) == 1 {
		return node
	}

	type item struct {
		r    Record
		dist float32
	}

	items := make([]item, len(records)-1)

	for i := 1; i < len(records); i++ {
		items[i-1] = item{
			r:    records[i],
			dist: Distance(node.Point.Vector, records[i].Vector),
		}
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].dist < items[j].dist
	})

	median := len(items) / 2
	node.Radius = items[median].dist

	left := make([]Record, 0, median)
	right := make([]Record, 0, len(items)-median)

	for i, it := range items {
		if i < median {
			left = append(left, it.r)
		} else {
			right = append(right, it.r)
		}
	}

	node.Left = Build(left)
	node.Right = Build(right)

	return node
}
