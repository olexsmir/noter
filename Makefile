.SILENT:
.PHONY:

build:
	go build -o .bin/api cmd/main.go

run: build
	.bin/api


migrate.new:
	migrate create -ext sql -dir migrations -seq $$name

export DBURI = postgres://postgres:qwerty@0.0.0.0:5432/noter?sslmode=disable

migrate.up:
	migrate -path migrations -database $$DBURI up

migrate.down:
	migrate -path migrations -database $$DBURI down 1

migrate.drop:
	migrate -path migrations -database $$DBURI drop
