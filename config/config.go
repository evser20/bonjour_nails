package config

import (
	"fmt"
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
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
