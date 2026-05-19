package main

import (
	"compress/gzip"
	"encoding/binary"
	"encoding/json"
	"log"
	"os"
)

type RawReference struct {
	Vector [14]float32 `json:"vector"`
	Label  string      `json:"label"`
}

func main() {
	inputFile, err := os.Open("resources/references.json.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	gzipReader, err := gzip.NewReader(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer gzipReader.Close()

	outputFile, err := os.Create("resources/references.bin")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	decoder := json.NewDecoder(gzipReader)

	_, err = decoder.Token()
	if err != nil {
		log.Fatal(err)
	}

	count := 0

	for decoder.More() {
		var raw RawReference

		if err := decoder.Decode(&raw); err != nil {
			log.Fatal(err)
		}

		err := binary.Write(outputFile, binary.LittleEndian, raw.Vector)
		if err != nil {
			log.Fatal(err)
		}

		var label uint8 = 0
		if raw.Label == "fraud" {
			label = 1
		}

		err = binary.Write(outputFile, binary.LittleEndian, label)
		if err != nil {
			log.Fatal(err)
		}

		count++

		if count%100000 == 0 {
			log.Printf("processed=%d", count)
		}
	}

	log.Printf("finished: %d vectors", count)
}
