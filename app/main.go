package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		input, err := reader.ReadString('\n')
		if err != nil {
			os.Exit(1)
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		args := strings.Fields(input)
		command := args[0]

		if isBuiltIn := HandleBuiltIn(command, args, input); isBuiltIn {
			continue
		}

		cmd := exec.Command(command, args[1:]...)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Printf("%s: command not found\n", command)
		}
	}
}

func HandleBuiltIn(command string, args []string, fullInput string) bool {
	switch command {
	case "exit":
		os.Exit(0)
		return true

	case "echo":
		if len(fullInput) > 5 {
			fmt.Println(fullInput[5:])
		} else {
			fmt.Println()
		}
		return true

	case "type":
		if len(args) < 2 {
			fmt.Println("type: missing argument")
			return true
		}
		argument := args[1]

		switch argument {
		case "echo", "type", "exit":
			fmt.Printf("%s is a shell builtin\n", argument)
		default:
			path, err := exec.LookPath(argument)
			if err != nil {
				fmt.Printf("%s: not found\n", argument)
			} else {
				fmt.Printf("%s is %s\n", argument, path)
			}
		}
		return true
	}

	return false
}
