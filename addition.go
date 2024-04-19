package main

import (
	"fmt"
)

func main() {
	const num1, num2 = 10, 20
	fmt.Println(num1 + num2)
	fmt.Print(num1 + num2)
	fmt.Printf("%d", num1+num2)
}
