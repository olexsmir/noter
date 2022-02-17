.SILENT:
.PHONY:

build:
	CGO_ENABLED=0 GOOS=linux go build -o .bin/api cmd/main.go	

run: build
	docker-compose up --remove-orphans api

test:
	GIN_MODE=release go test --short ./...

test.integration:
	GIN_MODE=release go test -v ./tests/

lint:
	golangci-lint run

swag:
	swag init -g cmd/main.go

migrate.new:
	migrate create -ext sql -dir migrations -seq $$name

export DBURI = postgres://postgres:qwerty@0.0.0.0:5432/noter?sslmode=disable

migrate.up:
	migrate -path migrations -database $$DBURI up

migrate.down:
	migrate -path migrations -database $$DBURI down 1

migrate.drop:
	migrate -path migrations -database $$DBURI drop

mockgen:
	mockgen -source=internal/service/service.go -destination=internal/service/mocks/mock.go
	mockgen -source=internal/repository/repository.go -destination=internal/repository/mocks/mock.go
