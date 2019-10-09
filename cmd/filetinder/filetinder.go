package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/shrotavre/filetinder/internal/filetinder"
	"github.com/shrotavre/filetinder/internal/server"
	"github.com/shrotavre/filetinder/internal/server/sdk"
	"github.com/shrotavre/filetinder/internal/shell"
)

var (
	subcommand string
	serverSDK  *sdk.FileTinderSDK
)

func init() {
	rawArgs := os.Args
	if len(rawArgs) >= 2 {
		subcommand = rawArgs[1]
	}

	// initialize SDK to talk to server
	serverSDK = sdk.New()
}

func main() {
	switch subcommand {
	case "start":
		binpath, err := filepath.Abs(os.Args[0])
		finErr(err)

		err = shell.ExecInBackground(binpath, "kickserver")
		finErr(err)

		fmt.Println("FileTinder started!")
		fmt.Printf("Open your http://localhost:%d to start choosing files", filetinder.Config.Port)
		return

	case "kickserver":
		fmt.Println("Running FileTinder server...")
		err := server.Start()
		finErr(err)
		break

	case "add":
		targetPath, err := filepath.Abs(os.Args[2])
		finErr(err)

		fstat, err := os.Stat(targetPath)
		finErr(err)

		switch mode := fstat.Mode(); {
		// If target is a folder
		case mode.IsDir():
			fpaths := make([]string, 0)
			err := filepath.Walk(targetPath, func(path string, info os.FileInfo, err error) error {
				if info.Mode().IsRegular() {
					fpaths = append(fpaths, path)
				}
				return nil
			})
			finErr(err)

			for _, fp := range fpaths {
				params := sdk.NewTargetParams{URL: fp}
				err = serverSDK.NewTarget(params)
				finErr(err)
			}

			fmt.Println("Added directory to FileTinder:", targetPath)
			break

		// If target is a single file
		case mode.IsRegular():
			params := sdk.NewTargetParams{URL: targetPath}
			err = serverSDK.NewTarget(params)
			finErr(err)

			fmt.Println("Added file to FileTinder:", targetPath)
			break
		}
		break

	case "remove":
		lsindexStr := os.Args[2]

		lsindex, err := strconv.Atoi(lsindexStr)
		finErr(err)

		ts, err := serverSDK.ListTarget()
		finErr(err)

		t := ts[lsindex-1]

		err = serverSDK.DelTarget(t.ID)
		finErr(err)

		fmt.Printf("File '%s' removed from dirtinder\n", t.URL)
		break

	case "list":
		ts, err := serverSDK.ListTarget()
		finErr(err)

		if len(ts) == 0 {
			fmt.Println("No files included yet.\nYou can add some files using 'filetinder add ./file/path'..")
			return
		}

		fmt.Println("List of included files:")
		for i, t := range ts {
			fmt.Printf("(%d) %s", i+1, t.URL)
			if len(t.Tags) > 0 {
				fmt.Print(" tagged: ", t.Tags)
			}

			fmt.Println()
		}
		break

	case "stop":
		err := serverSDK.KillServer()
		finErr(err)

		fmt.Println("FileTinder stopped...")
		break

	default:
		fmt.Println("Command unknown! To start FileTinder run 'filetinder start'.")
	}
}

func finErr(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(2)
	}
}
