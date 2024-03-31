PHONY: run
.SILENT:
run:
	docker-compose up -d --remove-orphans postgres
stop:
	docker-compose down -v
up:
	migrate -path internal/storage/postgres/migration -database 'postgresql://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up
down:
	migrate -path internal/storage/postgres/migration -database 'postgresql://postgres:qwerty@localhost:5432/postgres?sslmode=disable' down
go:
	go run ./cmd/app/main.go --config=./config/local.yaml
build:
	docker build -t app .  