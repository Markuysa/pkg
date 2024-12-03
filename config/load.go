package config

import (
	"os"

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

	return nil
}
