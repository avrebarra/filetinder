package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

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
			log.Fatal(err)
		}

		err = shell.ExecInBackground(binpath, "kickserver")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("FileTinder started!")
		fmt.Printf("Open your http://localhost:%d to start choosing files", config.DefaultPort)
		break

	case "kickserver":
		fmt.Println("Running FileTinder server...")
		if err := server.Start(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(2)
		}
		break

	case "add":
		targetPath, err := filepath.Abs(os.Args[2])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		message := map[string]interface{}{
			"url": targetPath,
		}

		bytePayload, err := json.Marshal(message)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		_, err = http.Post(fmt.Sprintf("http://localhost:%d/api/targets", config.DefaultPort), "application/json", bytes.NewBuffer(bytePayload))
		if err != nil {
			log.Fatalln(err)
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
