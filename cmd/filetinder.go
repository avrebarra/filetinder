package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/imroc/req"
	"github.com/shrotavre/filetinder/internal/config"
	"github.com/shrotavre/filetinder/internal/server"
	"github.com/shrotavre/filetinder/internal/shell"
)

var (
	subcommand string
)

func init() {
	rawArgs := os.Args
	if len(rawArgs) >= 2 {
		subcommand = rawArgs[1]
	}
}

func main() {
	switch subcommand {
	case "start":
		binpath, err := filepath.Abs(os.Args[0])
		if err != nil {
			handleErrorAndExit(err)
		}

		err = shell.ExecInBackground(binpath, "kickserver")
		if err != nil {
			handleErrorAndExit(err)
		}

		fmt.Println("FileTinder started!")
		fmt.Printf("Open your http://localhost:%d to start choosing files", config.DefaultPort)
		break

	case "kickserver":
		fmt.Println("Running FileTinder server...")
		if err := server.Start(); err != nil {
			handleErrorAndExit(err)
		}
		break

	case "add":
		targetPath, err := filepath.Abs(os.Args[2])
		if err != nil {
			handleErrorAndExit(err)
		}

		url := fmt.Sprintf("http://localhost:%d/api/targets", config.DefaultPort)
		payload := req.Param{
			"url": targetPath,
		}

		_, err = req.Post(url, req.Header{"Accept": "application/json"}, req.BodyJSON(payload))
		if err != nil {
			handleErrorAndExit(err)
		}

		fmt.Println("Added to FileTinder:", targetPath)
		break

	case "remove":
		break

	case "list":
		break

	case "stop":
		break

	default:
		fmt.Println("Command unknown! To start FileTinder run 'filetinder start'.")
	}
}

func handleErrorAndExit(err error) {
	fmt.Println("Error:", err)
	os.Exit(2)
}
