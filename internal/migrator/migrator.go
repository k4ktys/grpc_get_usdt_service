package migrator

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(migrationsPath, dbUser, dbPassword, dbPort, migrationsTable string) {
	m, err := migrate.New("file://"+migrationsPath,
		fmt.Sprintf("postgres://%s:%s@get_usdt_db:%s/%s?sslmode=disable", dbUser, dbPassword, dbPort, migrationsTable),
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
