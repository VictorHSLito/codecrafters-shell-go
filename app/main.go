package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Print("$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')

		command = command[:len(command)-1]

		if command == "exit" {
			os.Exit(0)
		}

		fmt.Println(command + ": not found")

		if err != nil {
			os.Exit(1)
		}
	}
}
