package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	fmt.Println("are u ready motherfucker")

	time.Sleep(2 * time.Second)
	fmt.Println("ok lets go")

	exePath, err := os.Executable()
	if err != nil {
		fmt.Printf("Error getting executable path: %v\n", err)
		return
	}

	exeDir := filepath.Dir(exePath)
	davirusExe := filepath.Join(exeDir, "davirus.exe")

	var cmd *exec.Cmd
	if _, err := os.Stat(davirusExe); err == nil {
		cmd = exec.Command("cmd.exe", "/c", "start", "cmd", "/k", davirusExe)
	} else {
		// Fallback for 'go run main.go' mode
		cmd = exec.Command("cmd.exe", "/c", "start", "cmd", "/k", "go", "run", "./davirus")
	}

	if err := cmd.Start(); err != nil {
		fmt.Printf("Failed to launch davirus: %v\n", err)
	}
}
