run:
	go run ./cmd/api

# DOCKER
docker-compose-build:
	docker compose build

docker-compose-up:
	docker compose up -d