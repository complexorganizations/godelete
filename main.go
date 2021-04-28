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
	findAllGoAppsBins()
	findAllGoAppsSource()
}

// Find all files in bin folder
func findAllGoAppsBins() {
	filepath.Walk(goUserBinPath, func(path string, info os.FileInfo, err error) error {
		if fileExists(path) {
			fileNameOnly := filepath.Base(path)
			fmt.Println(fileNameOnly)
			// take user input and if the user input matches the name than remove it.
		}
		return nil
	})
}

func findAllGoAppsSource() {
	filepath.Walk(goUserModPath, func(path string, info os.FileInfo, err error) error {
		if folderExists(path) {
			fmt.Println(path)
		}
		return nil
	})
}

func deleteBinAndSource(appname string) {
	//
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
