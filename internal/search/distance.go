package search

func Distance(a, b [14]float32) float32 {
	var sum float32

	for i := 0; i < 14; i++ {
		d := a[i] - b[i]
		sum += d * d
	}

	return sum
}
