package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

var (
	goUserPath    = fmt.Sprint(userDirectory() + "/go")
	goUserBinPath = fmt.Sprint(goUserPath + "/bin")
	goUserPkgPath = fmt.Sprint(goUserPath + "/pkg")
	goUserModPath = fmt.Sprint(goUserPkgPath + "/mod")
)

func init() {
	// System Requirements Check
	if !commandExists("go") {
		log.Fatal("Error: The application go was not found in the system.")
	}
	if !folderExists(goUserPath) {
		log.Fatal("Error: The go user folder was not found in the system.")
	}
}

func main() {
	findAllGoApps()
	takeUserInput()
}

func findAllGoApps() {
	// Find all files in bin folder
	filepath.Walk(goUserBinPath, func(path string, info os.FileInfo, err error) error {
		if fileExists(path) {
			fmt.Println(filepath.Base(path), fileSize(path))
		}
		return nil
	})
}

func takeUserInput() {
	// Take user input
	fmt.Println("Which package would you like to delete?")
	var appName string
	fmt.Scanln(&appName)
	if appName == "exit" {
		os.Exit(0)
	}
	deleteBinAndSource(appName)
}

func deleteBinAndSource(appname string) {
	// Remove the bins
	filepath.Walk(goUserBinPath, func(path string, info os.FileInfo, err error) error {
		if fileExists(path) {
			fileNameOnly := filepath.Base(path)
			if appname == fileNameOnly {
				os.Remove(path)
			}
		}
		return nil
	})
	// Remove the source
	filepath.Walk(goUserModPath, func(path string, info os.FileInfo, err error) error {
		if folderExists(path) {
			if strings.Contains(path, appname) {
				os.RemoveAll(path)
			}
		}
		return nil
	})
}

// Check if a folder exists
func folderExists(foldername string) bool {
	info, err := os.Stat(foldername)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// check if a file exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// Check if there is an app installed
func commandExists(cmd string) bool {
	cmd, err := exec.LookPath(cmd)
	if err != nil {
		return false
	}
	_ = cmd
	return true
}

// Get the current user dir
func userDirectory() string {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return user.HomeDir
}

// Get the size of a file
func fileSize(filepath string) int64 {
	file, err := os.Stat(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return file.Size()
}
