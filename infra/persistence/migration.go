package persistence

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/rellyson/car-sales-api/application/utils"
)

func ExecMigrations(db *sql.DB, tableName string) {
	logger := utils.NewLogger()
	driver, _ := postgres.WithInstance(db, &postgres.Config{
		MigrationsTable: "migrations_table",
	})

	m, err := migrate.NewWithDatabaseInstance(
		"file://./infra/persistence/migrations",
		tableName,
		driver)

	if err != nil {
		logger.Fatal(err.Error())
	}

	if err = m.Up(); err != nil && err.Error() != "no change" {
		logger.Fatal(err.Error())
	}

	logger.Info("[Persistence] - Database migrations syncronized succesfully!")
}
