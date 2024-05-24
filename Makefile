.PHONY: build
build:
	go build -v ./cmd/apiserver/

.PHONY: db
db-start:
	docker compose start
db-up:
	docker compose up -d
db-stop:
	docker compose stop
# db-psql:
# 	docker exec -it rest-api psql -U restapi restapi

.PHONY: test
test:
	go test -v -race -timeout 30s ./...



.DEFAULT_GOAL := build

