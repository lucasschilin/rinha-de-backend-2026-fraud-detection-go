package search

const K = 5

type Neighbor struct {
	Dist  float32
	Label uint8
}

func KNN(nodes []Node, records []Record, root int, target [14]float32, k int) []Neighbor {
	h := NewMaxHeap(k)

	var dfs func(i int)
	dfs = func(i int) {
		if i == -1 {
			return
		}

		n := nodes[i]

		d := Distance(target, records[n.Index].Vector)

		h.Push(Neighbor{
			Label: records[n.Index].Label,
			Dist:  d,
		})

		if n.Left != -1 {
			dfs(n.Left)
		}
		if n.Right != -1 {
			dfs(n.Right)
		}
	}

	dfs(root)

	return h.Items()
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

	return float64(fraud) / float64(n)
}
