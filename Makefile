run:
	go run ./cmd/api

# DOCKER
docker-compose-build:
	docker compose build

docker-compose-up:
	docker compose up -d

docker-compose-down:
	docker compose down

docker-api-image-build:	
	docker build --platform linux/amd64 -t ghcr.io/lucasschilin/rinha-de-backend-2026-fraud-detection-go:latest .

docker-api-image-run:
	docker run --rm -p 8080:8080 ghcr.io/lucasschilin/rinha-de-backend-2026-fraud-detection-go:latest

# PREPROCESS REFERENCES
preprocess-references:
	go run ./tools/preprocess/main.go