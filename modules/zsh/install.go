package zsh

import (
	"fmt"
	"os/exec"
)

func Install() {
	fmt.Println("Starting Zsh installation...")

	cmd := exec.Command("sudo", "apt", "install", "-y", "zsh")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error installing Zsh:", err)
		return
	}

	fmt.Println("Zsh installed successfully!")
}
