package database

import (
	"database/sql"
	"embed"
	"fmt"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	_ "modernc.org/sqlite"
)

//go:embed migrations
var migrationsFs embed.FS

// zerologMigrationLogger wraps a zerolog.Logger to make it compatible with golang-migrate.
type zerologMigrationLogger struct {
	logger zerolog.Logger
}

func (z zerologMigrationLogger) Printf(format string, v ...interface{}) {
	output := fmt.Sprintf(format, v...)
	// Printf normally requires explicitly adding a trailing newline, but Zerolog automatically inserts it for us
	output = strings.TrimSuffix(output, "\n")
	z.logger.Printf(output)
}

func (z zerologMigrationLogger) Verbose() bool {
	return z.logger.GetLevel() == zerolog.DebugLevel
}

var _ migrate.Logger = (*zerologMigrationLogger)(nil)

func NewMigrateInstance(connectionString string) (*migrate.Migrate, error) {
	db, err := sql.Open("sqlite", connectionString)
	if err != nil {
		return nil, fmt.Errorf("error opening db: %w", err)
	}

	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		return nil, fmt.Errorf("error creating migrate driver: %w", err)
	}

	migrationSource, err := iofs.New(migrationsFs, "migrations")
	if err != nil {
		return nil, fmt.Errorf("error creating migration source: %w", err)
	}

	migrateInstance, err := migrate.NewWithInstance(
		"iofs", migrationSource,
		"sqlite", driver,
	)
	if err != nil {
		return nil, err
	}

	migrateInstance.Log = zerologMigrationLogger{log.Logger}

	return migrateInstance, nil
}
