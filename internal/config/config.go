package config

import (
	"os"
	"strconv"
)

// AppConfigs object defining configuration used in whole application
type AppConfigs struct {
	Port int
}

// GetConfigs get cofiguration set for the whole application
func GetConfigs() *AppConfigs {
	d := &AppConfigs{}

	d.Port = 17763

	// If defined fetch from env
	if val, exists := os.LookupEnv("PORT"); exists {
		port, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}

		d.Port = port
	}

	return d
}
