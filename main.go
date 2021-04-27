package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
)

var (
	goUserPath    = fmt.Sprint(userDirectory() + "/go")
	goUserBinPath = fmt.Sprint(goUserPath + "/bin")
	goUserPkgPath = fmt.Sprint(goUserPath + "/pkg")
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
}

func findAllGoApps() {
	// Find all files in bin folder
	filepath.Walk(goUserBinPath, func(path string, info os.FileInfo, err error) error {
		if fileExists(path) {
			fileNameOnly := filepath.Base(path)
			fmt.Println(fileNameOnly)
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
	_ = cmd // variable declared and not used
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
