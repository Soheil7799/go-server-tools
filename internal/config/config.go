package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Server struct {
	Name        string `yaml:"name"`
	Host        string `yaml:"host"`
	Description string `yaml:"description"`
}

type SSHKey struct {
	Name        string `yaml:"name"`
	Path        string `yaml:"path"`
	Description string `yaml:"description"`
}

type Config struct {
	Servers []Server `yaml:"servers"`
	SSHKeys []SSHKey `yaml:"ssh_keys"`
}

func LoadConfig() (*Config, error) {
	// configPath := "~/.config/go-server-tools/config.yaml"
	configPath := "configs/example.yaml"

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
