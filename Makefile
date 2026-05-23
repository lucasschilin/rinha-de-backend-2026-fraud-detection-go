run:
	go run ./cmd/api

# DOCKER COMPOSE
docker-compose-build:
	docker compose build

docker-compose-up:
	docker compose up -d

docker-compose-down:
	docker compose down

# DOCKER IMAGE
docker-api-image-build:	
	docker build --platform linux/amd64 -t ghcr.io/lucasschilin/rinha-de-backend-2026-fraud-detection-go:latest .

docker-api-image-buildx-build:	
	docker buildx build --platform linux/amd64 -t ghcr.io/lucasschilin/rinha-de-backend-2026-fraud-detection-go:latest .

docker-api-image-run:
	docker run --rm -p 8080:9999 ghcr.io/lucasschilin/rinha-de-backend-2026-fraud-detection-go:latest

docker-api-image-push:
	docker push ghcr.io/lucasschilin/rinha-de-backend-2026-fraud-detection-go:latest

# PREPROCESS REFERENCES
preprocess-references:
	go run ./tools/preprocess/main.go