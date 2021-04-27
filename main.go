package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
)

var (
	goUserPath    = fmt.Sprint(userDirectory() + "/go")
	goUserBinPath = fmt.Sprint(goUserPath + "/bin")
)

func init() {
	// System Requirements Check
	if !commandExists("go") {
		log.Fatal("Error: The application go was not found in the system.")
	}
}

func main() {
	fmt.Println("Hello, playground")
}

func findAllGoApps() {
	files, err := os.ReadDir(goUserBinPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
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
