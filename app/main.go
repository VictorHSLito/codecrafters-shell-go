package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	for {
		fmt.Print("$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')

		command = command[:len(command)-1]

		if err != nil {
			os.Exit(1)
		}

		HandleCommand(command)
	}
}

func HandleCommand(command string) {
	prefix := strings.Split(command, " ")[0]

	switch prefix {

	case "exit":
		{
			os.Exit(0)
		}

	case "echo":
		{
			fmt.Println(command[5:])
		}

	case "type":
		{
			argument := strings.Split(command, " ")[1]
			switch argument {

			case "echo":
				{
					fmt.Printf("%s is a shell builtin\n", argument)
				}

			case "type":
				{
					fmt.Printf("%s is a shell builtin\n", argument)
				}

			case "exit":
				{
					fmt.Printf("%s is a shell builtin\n", argument)
				}
			default:
				path, err := exec.LookPath(argument)

				if err != nil {
					fmt.Printf("%s: not found\n", argument)
				}

				if len(path) > 0 {
					fmt.Printf("%s is %s\n", argument, path)
				}
			}
		}

	default:
		{
			fmt.Printf("%s: not found\n", command)
		}
	}
}
