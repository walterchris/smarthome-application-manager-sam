package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Name    string                   `yaml:"name,omitempty"`
	Plugins []map[string]interface{} `yaml:"plugins,omitempty"`
}

func Parse(Path string) (*Config, error) {
	file, err := os.ReadFile(Path)
	if err != nil {
		return nil, err
	}

	c := &Config{}

	err = yaml.Unmarshal(file, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
