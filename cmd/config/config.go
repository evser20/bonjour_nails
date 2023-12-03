package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App   `yaml:"app"`
		TgBot `yaml:"tg_bot"`
		HTTP  `yaml:"http"`
	}

	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}

	TgBot struct {
		Token string `yaml:"token"`
	}

	HTTP struct {
		Port string `yaml:"port"`
	}

	ServeConfig struct {
		PostgresAddress  string `toml:"postgres_address"`
		PostgresPort     uint   `toml:"postgres_port"`
		PostgresDBName   string `toml:"postgres_db_name"`
		PostgresUser     string `toml:"postgres_user"`
		PostgresPassword string `toml:"postgres_password"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./cmd/config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func Parse(c *ServeConfig) error {
	if _, err := toml.DecodeFile("./cmd/config/serve_config.toml", c); err != nil {
		return fmt.Errorf("config file parse error: %w", err)
	}
	return nil
}
