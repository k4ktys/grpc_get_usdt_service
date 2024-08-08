package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var migrationsPath, migrationsTable, dbUser, dbPassword, dbPort string

	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of migrations table")
	flag.StringVar(&dbUser, "db-user", "", "db user")
	flag.StringVar(&dbPassword, "db-password", "", "db password")
	flag.StringVar(&dbPort, "db-port", "", "db port")

	flag.Parse()

	if migrationsPath == "" {
		panic("migrations-path is required")
	}

	if dbUser == "" {
		panic("db-user is required")
	}

	if dbPassword == "" {
		panic("db-password is required")
	}

	if dbPort == "" {
		panic("db-port is required")
	}

	m, err := migrate.New("file://"+migrationsPath,
		fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", dbUser, dbPassword, dbPort, migrationsTable),
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations apply")

			return
		}

		panic(err)
	}

	fmt.Println("migrations applied successfully")
}
