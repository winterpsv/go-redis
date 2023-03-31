BINARY_NAME=main.out

build:
	go build -o build/${BINARY_NAME} ./cmd/app/main.go

run:
	go build -o build/${BINARY_NAME} ./cmd/app/main.go
	./build/${BINARY_NAME}

test:
	go test ./internal/adapter/db/mongodb/...
	go test ./internal/application/service/...
	go test ./internal/application/usecase/...

clean:
	go clean
	rm build/${BINARY_NAME}

start:
	docker-compose up

restart:
	docker-compose down
	sudo rm -rf mongodata/
	docker-compose build

stop:
	docker-compose down

cleandata:
	sudo rm -rf mongodata/