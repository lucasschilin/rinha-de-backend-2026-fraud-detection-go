package search

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

	if len(best) > k {
		best = best[:k]
	}

	return best
}

func Score(neighbors []Neighbor) float64 {
	var fraud int

	for n := 0; n < K; n++ {
		if neighbors[n].Label == 1 {
			fraud++
		}
	}

	return float64(fraud) / K
}
