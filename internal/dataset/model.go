package dataset

type rawReference struct {
	Vector [14]float32 `json:"vector"`
	Label  string      `json:"label"`
}
