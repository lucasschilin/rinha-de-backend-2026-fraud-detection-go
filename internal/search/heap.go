package search

type MaxHeap struct {
	data []Neighbor
	k    int
}

func NewMaxHeap(k int) *MaxHeap {
	return &MaxHeap{
		data: make([]Neighbor, 0, k),
		k:    k,
	}
}

func (h *MaxHeap) worst() float32 {
	if len(h.data) == 0 {
		return 1e30 // very large number
	}
	return h.data[0].Dist
}

func (h *MaxHeap) Push(n Neighbor) {
	// case 1: not full yet
	if len(h.data) < h.k {
		h.data = append(h.data, n)
		h.up(len(h.data) - 1)
		return
	}

	// case 2: worse than current best -> ignore
	if n.Dist >= h.worst() {
		return
	}

	// case 3: replace root safely
	h.data[0] = n
	h.down(0)
}

func (mh *MaxHeap) Items() []Neighbor {
	return mh.data
}

func (mh *MaxHeap) up(i int) {
	for i > 0 {
		p := (i - 1) / 2
		if mh.data[p].Dist >= mh.data[i].Dist {
			break
		}
		mh.data[p], mh.data[i] = mh.data[i], mh.data[p]
		i = p
	}
}

func (mh *MaxHeap) down(i int) {
	n := len(mh.data)
	if n == 0 {
		return
	}

	for {
		l := 2*i + 1
		r := 2*i + 2
		largest := i

		if l < n && mh.data[l].Dist > mh.data[largest].Dist {
			largest = l
		}
		if r < n && mh.data[r].Dist > mh.data[largest].Dist {
			largest = r
		}
		if largest == i {
			break
		}
		mh.data[i], mh.data[largest] = mh.data[largest], mh.data[i]
		i = largest
	}
}
