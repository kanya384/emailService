package config

import (
	"fmt"

	"github.com/vrischmann/envconfig"
)

type Config struct {
	App struct {
		Host string
		Port int
	}

	Log struct {
		Level string
	}

	Pg struct {
		User    string
		Pass    string
		DbName  string
		Host    string
		Port    string
		PoolMax int
	}

	FileStore struct {
		Path string
	}

	Email struct {
		Host  string
		Port  int
		Login string
		Pass  string
	}
}

func InitConfig(prefix string) (*Config, error) {
	conf := &Config{}
	if err := envconfig.InitWithPrefix(conf, prefix); err != nil {
		return nil, fmt.Errorf("init config error: %w", err)
	}

	return conf, nil
}
