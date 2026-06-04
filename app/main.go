package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("$ ")
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Printf("%s: command not found", command)

	if err != nil {
		os.Exit(1)
	}
}
