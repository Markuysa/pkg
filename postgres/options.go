package postgres

type options struct {
	migrate *MigrateCfg
}

type option interface {
	apply(*options)
}

type migrateOptions struct {
	migrate *MigrateCfg
}

func (m *migrateOptions) apply(opts *options) {
	opts.migrate = m.migrate
}

func WithMigrate(m *MigrateCfg) option {
	return &migrateOptions{migrate: m}
}
