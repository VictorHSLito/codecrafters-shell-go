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

		fmt.Println(command[:len(command)-1] + ": not found")

		if err != nil {
			os.Exit(1)
		}
	}
}
