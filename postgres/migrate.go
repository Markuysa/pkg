package postgres

import (
	"embed"
	"fmt"

	"github.com/Markuysa/pkg/log"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type MigrateCfg struct {
	MigratePath string `default:"."`
	Fs          embed.FS
	PgDSN       string
}

func applyMigrations(
	cfg *MigrateCfg,
) error {
	d, err := iofs.New(cfg.Fs, cfg.MigratePath)
	if err != nil {
		return errors.Wrap(err, "embed postgres migrations")
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, cfg.PgDSN)
	if err != nil {
		return errors.Wrap(err, "apply postgres migrations")
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("failed to apply migrations: %v", err)
	}

	version, dirty, err := m.Version()
	if err != nil {
		return fmt.Errorf("failed to get migration version: %v", err)
	}

	if dirty {
		return fmt.Errorf("dirty migration version: %d", version)
	}

	log.Infof("migrations successfully applied with version %d", version)

	return nil
}
