MAIN_PKG_PATH := ./cmd/main/main.go

migrate:
	go run ./cmd/migrator --migrations-path=./migrations --db-user=postgres --db-password=simplePassword --db-port=5432 --migrations-table=postgres

run:
	go run ${MAIN_PKG_PATH}