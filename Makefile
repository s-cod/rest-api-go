.PHONY: build
build:
	go build -v ./cmd/apiserver/
run:
	go run ./cmd/apiserver/

.PHONY: db
db-stop:
	docker compose start
db:
	docker compose up -d
# run:
# 	docker run --name api-go -e POSTGRES_DB=rest-api -e POSTGRES_USER=restapi -e POSTGRES_PASSWORD=restapi -v "${pwd}\database:/var/lib/postgresql/data" -d postgres
# db-psql:
# docker exec -it 5d8583c47e4e psql -U restapi  rest-api

db-stop:
	docker compose stop

.PHONY: test
test:
	go test -v -race -timeout 30s ./...



.DEFAULT_GOAL := build

