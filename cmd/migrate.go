package cmd

import (
	"fmt"

	"github.com/Ow1Dev/QustionOfToday/internal/database"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
	"github.com/urfave/cli/v2"
)

var CmdMigrate = &cli.Command{
	Name:   "migrate",
	Usage:  "Migrate the database",
	Action: runMigrate,
}

func runMigrate(ctx *cli.Context) error {
	db, err := database.Connect(ctx.Context)
	if err != nil {
		return err
	}

	sqlDB := stdlib.OpenDBFromPool(db) // Convert pgxpool.Pool to *sql.DB
	if sqlDB == nil {
		return fmt.Errorf("failed to convert pgxpool.Pool to *sql.DB")
	}
	defer sqlDB.Close()

	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("failed to create migrator: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration failed: %w", err)
	}

	return nil
}
