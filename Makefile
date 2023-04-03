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

buildimage:
	docker build -t task3_4_image .

clean:
	go clean
	rm build/${BINARY_NAME}

start:
	docker-compose  up

restart:
	docker-compose down
	docker build -t task3_4_image .
	docker-compose up

stop:
	docker-compose down