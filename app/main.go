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

		if command == "exit" {
			os.Exit(0)
		}

		if strings.HasPrefix(command, "echo ") {
			fmt.Println(command[5:])
		} else {
			fmt.Printf("%s: not found\n", command)
		}

		if err != nil {
			os.Exit(1)
		}
	}
}
