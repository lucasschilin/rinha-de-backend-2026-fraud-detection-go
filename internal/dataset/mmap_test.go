package dataset

import (
	"encoding/binary"
	"log"
	"math"
	"testing"

	"golang.org/x/exp/mmap"
)

const (
	floatCount = 14
	floatSize  = 4
	recordSize = floatCount*floatSize + 1
)

func TestMap(t *testing.T) {
	reader, err := mmap.Open("../../resources/references.bin")
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	defer reader.Close()

	buf := make([]byte, recordSize)

	_, err = reader.ReadAt(buf, 0)
	if err != nil {
		t.Fatalf("failed to read file: %v", err)
	}

	for i := 0; i < floatCount; i++ {
		offset := i * floatSize

		bits := binary.LittleEndian.Uint32(
			buf[offset : offset+floatSize],
		)

		value := math.Float32frombits(bits)

		log.Printf("float[%d] = %f", i, value)
	}

	label := uint8(buf[56])

	log.Printf("label = %d", label)
}
