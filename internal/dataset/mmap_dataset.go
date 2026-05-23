package dataset

import (
	"encoding/binary"
	"math"

	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/search"
	"golang.org/x/exp/mmap"
)

const (
	vectorSize = 14 * 4
	labelSize  = 1
	recordSize = vectorSize + labelSize
)

type MmapDataset struct {
	Reader *mmap.ReaderAt
	Count  int
}

func LoadMmap(path string, count int) (*MmapDataset, error) {
	reader, err := mmap.Open(path)
	if err != nil {
		return nil, err
	}

	return &MmapDataset{
		Reader: reader,
		Count:  count,
	}, nil
}

func (d *MmapDataset) LoadAll() ([]search.Record, error) {
	records := make([]search.Record, d.Count)

	buf := make([]byte, recordSize)

	for i := 0; i < d.Count; i++ {
		offset := int64(i * recordSize)

		_, err := d.Reader.ReadAt(buf, offset)
		if err != nil {
			return nil, err
		}

		var v [14]float32

		for j := 0; j < 14; j++ {
			bits := binary.LittleEndian.Uint32(buf[j*4 : (j+1)*4])
			v[j] = math.Float32frombits(bits)
		}

		records[i] = search.Record{
			Vector: v,
			Label:  buf[56],
		}
	}

	return records, nil
}
