package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/imroc/req"
	"github.com/shrotavre/filetinder/internal/config"
	"github.com/shrotavre/filetinder/internal/filetinder"
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
		strid := os.Args[2]

		id, err := strconv.Atoi(strid)
		if err != nil {
			handleErrorAndExit(err)
		}

		url := fmt.Sprintf("http://localhost:%d/api/targets/%d", config.DefaultPort, id)
		_, err = req.Delete(url, req.Header{"Accept": "application/json"})
		if err != nil {
			handleErrorAndExit(err)
		}

		fmt.Printf("File with id:%d removed from dirtinder\n", id)
		break

	case "list":
		url := fmt.Sprintf("http://localhost:%d/api/targets", config.DefaultPort)
		r, err := req.Get(url, req.Header{"Accept": "application/json"})
		if err != nil {
			handleErrorAndExit(err)
		}

		ts := filetinder.TargetStoreInst.List()
		r.ToJSON(&ts)

		fmt.Println("List of included files:")
		for i, t := range ts {
			fmt.Printf("(%d) id:%d : %s", i+1, t.ID, t.URL)
			if len(t.Tags) > 0 {
				fmt.Print(" tagged: ", t.Tags)
			}

			fmt.Println()
		}

		break

	case "stop":
		url := fmt.Sprintf("http://localhost:%d/api/funcs/stop-server", config.DefaultPort)

		_, err := req.Post(url)
		if err != nil {
			handleErrorAndExit(err)
		}

		fmt.Println("FileTinder stopped...")
		break

	default:
		fmt.Println("Command unknown! To start FileTinder run 'filetinder start'.")
	}
}

func handleErrorAndExit(err error) {
	fmt.Println("Error:", err)
	os.Exit(2)
}
