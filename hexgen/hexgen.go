package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime/debug"
)

var (
	version = ""
	commit  = "none"
	date    = "unknown"
)

var (
	destination = flag.String("path", "", "New folder name")
	mod         = flag.String("mod", "", "New mod name")
	showVersion = flag.Bool("version", false, "Print version.")
)

const usageText = `hexgen to generate hexagonal architecture folders`

func usage() {
	_, _ = io.WriteString(os.Stderr, usageText)
	flag.PrintDefaults()
}

func printVersion() {
	if version != "" {
		fmt.Printf("v%s\nCommit: %s\nDate: %s\n", version, commit, date)
	} else {
		printModuleVersion()
	}
}

func printModuleVersion() {
	if bi, exists := debug.ReadBuildInfo(); exists {
		fmt.Println(bi.Main.Version)
	} else {
		log.Printf("No version information found. Make sure to use " +
			"GO111MODULE=on when running 'go get' in order to use specific " +
			"version of the binary.")
	}
}

// createDefaultFolders creates the folders and files for the hexagonal architecture
func createDefaultFolders(path string) (err error) {
	// Create the root folder
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return
	}

	// Create the internal folder
	err = os.MkdirAll(fmt.Sprintf("%v%v", path, "/internal"), os.ModePerm)
	if err != nil {
		return
	}

	// Create the domain folder
	err = os.MkdirAll(fmt.Sprintf("%v%v", path, "/internal/domain"), os.ModePerm)
	if err != nil {
		return
	}

	// Create the service folder
	err = os.MkdirAll(fmt.Sprintf("%v%v", path, "/internal/domain/service"), os.ModePerm)
	if err != nil {
		return
	}

	// Create the port folder
	err = os.MkdirAll(fmt.Sprintf("%v%v", path, "/internal/domain/port"), os.ModePerm)
	if err != nil {
		return
	}

	// Create the port service folder
	err = os.MkdirAll(fmt.Sprintf("%v%v", path, "/internal/domain/port/service"), os.ModePerm)
	if err != nil {
		return
	}

	// Create the adapter folder
	err = os.MkdirAll(fmt.Sprintf("%v%v", path, "/internal/adapter"), os.ModePerm)
	if err != nil {
		return
	}

	// Create the handler folder
	err = os.MkdirAll(fmt.Sprintf("%v%v", path, "/internal/handler"), os.ModePerm)
	if err != nil {
		return
	}

	// Create the http folder
	err = os.MkdirAll(fmt.Sprintf("%v%v", path, "/internal/http"), os.ModePerm)
	if err != nil {
		return
	}

	// Create the config folder
	err = os.MkdirAll(fmt.Sprintf("%v%v", path, "/config"), os.ModePerm)
	if err != nil {
		return
	}

	return nil
}

// createDefaultFolders creates the folders and files for the hexagonal architecture
func creatDefaultFiles(path string) (err error) {
	// Create main.go
	err = os.WriteFile(fmt.Sprintf("%v%v", path, "/main.go"), []byte(`package main`), 0644)
	if err != nil {
		return
	}

	// Create config.json
	err = os.WriteFile(fmt.Sprintf("%v%v", path, "/config.json"), nil, 0644)
	if err != nil {
		return
	}

	// Create the config.go file
	err = os.WriteFile(fmt.Sprintf("%v%v", path, "/config/config.go"), []byte(`package config`), 0644)
	if err != nil {
		return
	}

	// Create the port service.go
	err = os.WriteFile(fmt.Sprintf("%v%v", path, "/internal/domain/port/service/service.go"), []byte(`package service`), 0644)
	if err != nil {
		return
	}

	// Create the server folder
	err = os.WriteFile(fmt.Sprintf("%v%v", path, "/internal/http/server.go"), []byte(`package http`), 0644)
	if err != nil {
		return
	}

	// Create the router folder
	err = os.WriteFile(fmt.Sprintf("%v%v", path, "/internal/http/router.go"), []byte(`package http`), 0644)
	if err != nil {
		return
	}

	// Create the handler folder
	err = os.WriteFile(fmt.Sprintf("%v%v", path, "/internal/http/handlers.go"), []byte(`package http`), 0644)
	if err != nil {
		return
	}

	return nil
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if *showVersion {
		printVersion()
		return
	}

	// Check if the folder path is provided
	if *destination == "" {
		fmt.Println("Please provide a folder path using the -path flag.")
		return
	}

	// Create default folders for hex
	err := createDefaultFolders(*destination)
	if err != nil {
		fmt.Printf("Error creating folder: %v\n", err)
		return
	} else {
		fmt.Println("Folder created successfully")
	}

	// Create default files for hex
	err = creatDefaultFiles(*destination)
	if err != nil {
		fmt.Printf("Error creating files: %v\n", err)
		return
	} else {
		fmt.Println("Files created successfully")
	}

	// Run go mod init if mod is provided
	if *mod != "" {
		cmd := exec.Command("go", "mod", "init", *mod)
		cmd.Dir = *destination
		output, errMod := cmd.CombinedOutput()
		if errMod != nil {
			fmt.Printf("failed to run 'go mod init': %v, output: %s", errMod, string(output))
			return
		} else {
			fmt.Println("go mod init completed successfully")
		}
	}
}
