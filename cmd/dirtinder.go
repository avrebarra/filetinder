package main

import (
	"fmt"
	"os"

	"github.com/shrotavre/filetinder/internal/config"
	"github.com/shrotavre/filetinder/internal/server"
	"github.com/shrotavre/filetinder/internal/shell"
)

func main() {
	rawArgs := os.Args
	if len(rawArgs) < 2 {
		fmt.Println("Command not supplied! To start dirtinder run 'dirtinder start'.")
		return
	}

	subcommand := rawArgs[1]

	switch subcommand {
	case "start":
		appconf := config.GetConfigs()

		shell.ExecInBackground("./dirtinder kickserver")

		fmt.Println("Dirtinder started!")
		fmt.Printf("Open your http://localhost:%s to start choosing files", appconf.Port)
		break

	case "kickserver":
		server.Start()
		break

	case "add":
		break

	case "remove":
		break

	case "stop":
		break

	default:
		fmt.Println("Command unknown!")
	}
}
