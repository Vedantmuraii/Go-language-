package main

import "fmt"

func main() {
	result := display()
	fmt.Println("welcome to: ", *result)
}

func display() *string {
	message := "SRKNEC"
	return &message
}
