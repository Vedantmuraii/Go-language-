package main

import (
	"fmt"
	"os"
)

func main() {
	programName := os.Args[0]
	fmt.Println("total arguments: ", len(os.Args))
	fmt.Println("program name: ", programName)

	fmt.Println("\nArguments: ")
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("\tArguments[%d]: %s\n", i, os.Args[i])
	}
}
