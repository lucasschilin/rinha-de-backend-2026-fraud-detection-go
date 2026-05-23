package search

func Build(records []Record) ([]Node, int) {
	nodes := make([]Node, len(records))

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

		nodes[root].Radius = Distance(
			records[root].Vector,
			records[mid].Vector,
		)

		nodes[root].Left = build(start+1, mid)
		nodes[root].Right = build(mid, end)

		return root
	}

	root := build(0, len(records))

	return nodes, root
}
