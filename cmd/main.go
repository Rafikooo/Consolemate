package main

import (
	"Consolemate/modules/fish"
	"Consolemate/modules/zsh"
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "init":
			initShell()
		default:
			fmt.Printf("Unknown command: %s\n", command)
			showHelp()
		}
	} else {
		showHelp()
	}
}

func initShell() {
	fmt.Println("Welcome to Consolemate!")
	fmt.Println("Which shell would you like to install?")
	fmt.Println("1) Fish")
	fmt.Println("2) Zsh")
	fmt.Println("3) Cancel")

	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')

	switch choice {
	case "1\n":
		fmt.Println("Installing Fish...")
		fish.Install()
	case "2\n":
		fmt.Println("Installing Zsh...")
		zsh.Install()
	case "3\n":
		fmt.Println("Canceled.")
	default:
		fmt.Println("Unknown option, please try again.")
	}
}

func showHelp() {
	fmt.Println("Consolemate - Your system management tool")
	fmt.Println("Usage:")
	fmt.Println("  cm [command]")
	fmt.Println("Available commands:")
	fmt.Println("  init       - Initialize tool (choose a shell to install)")
	fmt.Println("  help       - Display this summary")
}
