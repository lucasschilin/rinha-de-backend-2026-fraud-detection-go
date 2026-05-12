package main

import (
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/config"
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/dataset"
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/router"
)

func main() {
	cfg := config.Load()

	start := time.Now()

	ds, err := dataset.Load("resources/references.json.gz")
	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	log.Printf(
		"dataset loaded: vectors=%d labels=%d startup=%s memory=%.2fMB",
		len(ds.Vectors),
		len(ds.Labels),
		elapsed,
		float64(mem.Alloc)/1024/1024,
	)

	r := router.New(ds)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	log.Printf("API listening on :%s", cfg.Port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
