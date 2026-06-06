package main

import (
	"bufio"
	"fmt"
	"os"
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
			os.Exit(1)
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
				{
					fmt.Printf("%s: not found\n", argument)
				}
			}
		}

	default:
		{
			fmt.Printf("%s: not found\n", command)
		}
	}
}
