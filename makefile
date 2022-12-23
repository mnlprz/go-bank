build:
	@go build -o bin/go-bank

run: build
	@./bin/go-bank

test:
	@go test -v ./...

postgres:
	sudo docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine

createdb:
	sudo docker exec -it postgres createdb --username=root --owner=root go-bank

dropdb:
	docker exec -it postgres dropdb go-bank
