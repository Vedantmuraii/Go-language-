package main

import "fmt"

func update(num *int) {
	*num = 30
}
func main() {
	var number = 55
	update(&number)
	fmt.Println("the number is: ", number)
}
