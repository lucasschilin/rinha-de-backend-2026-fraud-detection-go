package dataset

import "github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/vector"

type Dataset struct {
	Vectors []vector.Vector
	Labels  []uint8 // 0 = non-fraud, 1 = fraud
}
