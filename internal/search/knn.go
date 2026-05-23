package search

const K = 5

type Neighbor struct {
	Dist  float32
	Label uint8
}

func KNN(nodes []Node, records []Record, root int, target [14]float32, k int) []Neighbor {
	h := NewMaxHeap(k)

	var search func(i int)
	search = func(i int) {
		if i == -1 {
			return
		}

		n := nodes[i]

		d := Distance(target, records[n.Index].Vector)

		h.Push(Neighbor{
			Label: records[n.Index].Label,
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
			worst := h.worst()

			if d-n.Radius <= worst && d+n.Radius >= 0 {
				search(second)
			}
		}
	}

	search(root)

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
