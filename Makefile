MAIN_PKG_PATH := ./cmd/main/main.go
LOCAL_CONFIG_PATH := ./config/local.yaml
BINARY_NAME := get_usdt_grpc_service

run:
	go run ${MAIN_PKG_PATH} --config=${LOCAL_CONFIG_PATH}