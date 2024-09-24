package fish

import (
	"fmt"
	"os/exec"
)

func Install() {
	fmt.Println("Starting Fish installation...")

	cmd := exec.Command("sudo", "apt-add-repository", "ppa:fish-shell/release-3")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error adding Fish repository:", err)
		return
	}

	cmd = exec.Command("sudo", "apt", "update")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error updating package list:", err)
		return
	}

	cmd = exec.Command("sudo", "apt", "install", "-y", "fish")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error installing Fish:", err)
		return
	}

	fmt.Println("Fish installed successfully!")
}
