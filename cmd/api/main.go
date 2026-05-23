package main

import (
	"log"
	"net/http"
	"runtime"
	"time"

	_ "net/http/pprof"

	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/config"
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/dataset"
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/router"
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/search"
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/service"
	"github.com/lucasschilin/rinha-de-backend-2026-fraud-detection-go/internal/vector"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	cfg := config.Load()

	start := time.Now()

	ds, err := dataset.LoadMmap("resources/references.bin", 3_000_000)
	if err != nil {
		log.Fatal(err)
	}

	records, err := ds.LoadAll()
	if err != nil {
		log.Fatal(err)
	}
	tree := search.Build(records)

	elapsed := time.Since(start)

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	log.Printf(
		"dataset loaded: startup=%s memory=%.2fMB",
		elapsed,
		float64(mem.Alloc)/1024/1024,
	)

	log.Println(ds)

	vectorBuilder := vector.NewBuilder()

	fraudService := service.NewFraudService(vectorBuilder, tree)
	r := router.New(fraudService)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	log.Printf("API listening on :%s", cfg.Port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
