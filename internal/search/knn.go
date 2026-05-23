package search

const K = 5

type Neighbor struct {
	Dist  float32
	Label uint8
}

func KNN(root *Node, target [14]float32, k int) []Neighbor {
	h := NewMaxHeap(k)

	var search func(n *Node)

	search = func(n *Node) {
		if n == nil {
			return
		}

		d := Distance(target, n.Point.Vector)

		h.Push(Neighbor{
			Label: n.Point.Label,
			Dist:  d,
		})

		if n.Left == nil && n.Right == nil {
			return
		}

		if d < n.Radius {
			search(n.Left)

			if d+n.Radius >= h.worst() {
				search(n.Right)
			}
		} else {
			search(n.Right)

			if d-n.Radius <= h.worst() {
				search(n.Left)
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
