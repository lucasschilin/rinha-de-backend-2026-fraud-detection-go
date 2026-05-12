package dataset

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"os"

	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/vector"
)

func Load(path string) (*Dataset, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return nil, fmt.Errorf("create gzip reader: %w", err)
	}
	defer gzipReader.Close()

	decoder := json.NewDecoder(gzipReader)

	dataset := &Dataset{
		Vectors: make([]vector.Vector, 0, 3_000_000),
		Labels:  make([]uint8, 0, 3_000_000),
	}

	_, err = decoder.Token()
	if err != nil {
		return nil, fmt.Errorf("read opening token: %w", err)
	}

	for decoder.More() {
		var raw rawReference

		if err := decoder.Decode(&raw); err != nil {
			return nil, fmt.Errorf("decode reference: %w", err)
		}

		dataset.Vectors = append(dataset.Vectors, raw.Vector)

		if raw.Label == "fraud" {
			dataset.Labels = append(dataset.Labels, 1)
		} else {
			dataset.Labels = append(dataset.Labels, 0)
		}
	}

	return dataset, nil

}
