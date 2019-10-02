package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/shrotavre/filetinder/internal/config"
	"github.com/shrotavre/filetinder/internal/server"
	"github.com/shrotavre/filetinder/internal/shell"
)

var (
	subcommand string
	binpath    string
)

func init() {
	rawArgs := os.Args
	if len(rawArgs) >= 2 {
		subcommand = rawArgs[1]
	}

	v, err := filepath.Abs(os.Args[0])
	if err != nil {
		log.Fatal(err)
	}

	binpath = v
}

func main() {
	switch subcommand {
	case "start":
		port := config.DefaultPort

		err := shell.ExecInBackground(binpath, "kickserver")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("FileTinder started!")
		fmt.Printf("Open your http://localhost:%d to start choosing files", port)
		break

	case "kickserver":
		fmt.Println("Running FileTinder server...")
		if err := server.Start(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(2)
		}
		break

	case "add":
		break

	case "remove":
		break

	case "stop":
		break

	default:
		fmt.Println("Command unknown! To start FileTinder run 'filetinder start'.")
	}
}
