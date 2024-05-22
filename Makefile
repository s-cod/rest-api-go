.PHONY: build
build:
	go build -v ./cmd/apiserver/

.PHONY: db
db-start:
	docker start habr-pg
db-run:
	docker run --name habr-pg -p 5432:5432 -e POSTGRES_USER=s -e POSTGRES_PASSWORD=s -e POSTGRES_DB=restapi_dev -d postgres
db-psql:
	docker exec -it habr-pg psql -U sofron restapi_dev

.PHONY: test
test:
	go test -v -race -timeout 30s ./...



.DEFAULT_GOAL := build

