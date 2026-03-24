package main

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Hosts []Host `toml:"hosts"`
}

type Bastion struct {
	Name string `toml:"name"`
	Host []Host
}

type Host struct {
	Name    string `toml:"name"`
	Address string `toml:"address"`
	User    string `toml:"user"`
	Key     string `toml:"key,omitempty"`
}

func loadConfig(path string) (Config, error) {
	var cfg Config
	_, err := toml.DecodeFile(path, &cfg)
	return cfg, err
}

func saveConfig(path string, cfg Config) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return toml.NewEncoder(f).Encode(cfg)
}
