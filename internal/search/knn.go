package search

import "sort"

const K = 5

type Neighbor struct {
	Dist  float32
	Label uint8
}

func KNN(root *Node, target [14]float32, k int) []Neighbor {
	best := make([]Neighbor, 0, k)

	var dfs func(n *Node)
	dfs = func(n *Node) {
		if n == nil {
			return
		}

		d := Distance(target, n.Point.Vector)
		best = append(best, Neighbor{
			Label: n.Point.Label,
			Dist:  d,
		})

		dfs(n.Left)
		dfs(n.Right)
	}

	dfs(root)

	sort.Slice(best, func(i, j int) bool {
		return best[i].Dist < best[j].Dist
	})

	if len(best) > k {
		best = best[:k]
	}

	return best
}

func Score(neighbors []Neighbor) float64 {
	var fraud int

	n := K
	if len(neighbors) < n {
		n = len(neighbors)
	}

	for i := 0; i < n; i++ {
		if neighbors[i].Label == 1 {
			fraud++
		}
	}

	return float64(fraud) / float64(K)
}
