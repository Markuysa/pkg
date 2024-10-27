package migrate

import "embed"

func Migrate(
	fs embed.FS,
	migPath string,
	pgDSN string,
) (*migrate.Migrate, error) {
	d, err := iofs.New(fs, migPath)
	if err != nil {
		return nil, errors.Wrap(err, "embed postgres migrations")
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, pgDSN)
	if err != nil {
		return nil, errors.Wrap(err, "apply postgres migrations")
	}

	return m, nil
}
