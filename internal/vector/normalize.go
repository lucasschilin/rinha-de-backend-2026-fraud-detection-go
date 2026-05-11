package vector

func Clamp(value float64) float32 {
	if value < 0 {
		return 0
	}

	if value > 1 {
		return 1
	}

	return float32(value)
}
