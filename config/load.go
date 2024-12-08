package config

import (
	"os"

	"github.com/creasty/defaults"
	"github.com/go-playground/validator"
	"gopkg.in/yaml.v2"
)

// TODO add vault.
func LoadFromYAML(cfg any, path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(file, cfg); err != nil {
		return err
	}

	err = defaults.Set(cfg)
	if err != nil {
		return err
	}

	err = validator.New().Struct(cfg)
	if err != nil {
		return err
	}

	return nil
}
