MAIN_PKG_PATH := ./cmd/main/main.go

run:
	go run ${MAIN_PKG_PATH}

lint:
	golangci-lint run

build:
	go build -o main ${MAIN_PKG_PATH} 