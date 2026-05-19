package dataset

import (
	"encoding/binary"
	"io"
	"os"

	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/vector"
)

func LoadBinary(path string) (*Dataset, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	ds := &Dataset{
		Vectors: make([]vector.Vector, 0, 3_000_000),
		Labels:  make([]uint8, 0, 3_000_000),
	}

	for {
		var vec vector.Vector
		var label uint8

		err := binary.Read(file, binary.LittleEndian, &vec)
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		err = binary.Read(
			file,
			binary.LittleEndian,
			&label,
		)

		if err != nil {
			return nil, err
		}

		ds.Vectors = append(ds.Vectors, vec)
		ds.Labels = append(ds.Labels, label)
	}

	return ds, nil
}
