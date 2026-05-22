package dataset

import (
	"encoding/binary"
	"math"

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

func (d *MmapDataset) GetRecord(index int) ([14]float32, uint8, error) {
	var vector [14]float32

	offset := int64(index * recordSize)

	buf := make([]byte, recordSize)

	_, err := d.Reader.ReadAt(buf, offset)
	if err != nil {
		return vector, 0, err
	}

	for i := 0; i < 14; i++ {
		bits := binary.LittleEndian.Uint32(
			buf[i*4 : (i+1)*4],
		)
		vector[i] = math.Float32frombits(bits)
	}
	label := buf[56]

	return vector, label, nil
}
