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
		fmt.Println("Command not supplied! To start FileTinder run 'filetinder start'.")
		return
	}

	subcommand := rawArgs[1]

	switch subcommand {
	case "start":
		appconf := config.GetConfigs()

		err := shell.ExecInBackground("./filetinder kickserver")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("FileTinder started!")
		fmt.Printf("Open your http://localhost:%d to start choosing files", appconf.Port)
		break

	case "kickserver":
		fmt.Println("Running FileTinder server...")
		if err := server.Start(); err != nil {
			fmt.Println("Error:", err)
			return
		}

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
