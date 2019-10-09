package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
)

const (
	defaultPort int = 17763
)

// FileTinderConfig holds all configurations used in app
type FileTinderConfig struct {
	Env     string
	Port    int
	WorkDir string
	UIPath  string
}

func getEnvWithFallback(key string, def string) string {
	v := os.Getenv(key)
	if v == "" {
		v = def
	}

	return v
}

func panicOnErr(err error) {
	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}
}

// InitConfig initializes config from files or default values
func InitConfig() *FileTinderConfig {
	var err error

	binpath, _ := filepath.Abs(os.Args[0])
	pwd := filepath.Dir(binpath)

	// Create with default values
	conf := FileTinderConfig{
		Env:     "production",
		Port:    defaultPort,
		WorkDir: ".",
		UIPath:  "./ui/public",
	}

	conf.Env = getEnvWithFallback("ENV", conf.Env)
	port, _ := strconv.Atoi(getEnvWithFallback("PORT", "-1"))

	if conf.Env == "production" {
		conf.WorkDir = pwd
		conf.UIPath = filepath.Join(pwd, "./ui")
	}

	if port != -1 {
		conf.Port = port
	}

	conf.WorkDir, err = filepath.Abs(conf.WorkDir)
	panicOnErr(err)

	conf.UIPath, err = filepath.Abs(conf.UIPath)
	panicOnErr(err)

	return &conf
}
