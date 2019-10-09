package config

import (
	"os"
	"path/filepath"
)

const (
	defaultPort int = 17763
)

// FileTinderConfig holds all configurations used in app
type FileTinderConfig struct {
	Port    int
	WorkDir string
}

// InitConfig initializes config from files or default values
func InitConfig() *FileTinderConfig {
	wd, err := filepath.Abs(os.Args[0])
	if err != nil {
		panic(err)
	}

	c := FileTinderConfig{
		Port:    defaultPort,
		WorkDir: wd,
	}

	return &c
}
