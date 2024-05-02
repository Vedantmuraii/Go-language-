package main

import "fmt"

func main() {
	var name = "john"
	var ptr *string

	ptr = &name
	fmt.Println("value of pointer is: ", ptr)
	fmt.Println("address of the variable: ", &name)

	fmt.Println("value of pointer is: ", *ptr)
	fmt.Println("address of the variable: ", name)
}
