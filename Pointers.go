package main

import "fmt"

func main() {
	var intData = 20
	var intPointer *int
	intPointer = &intData
	fmt.Println("what intData stores: ", intData)
	fmt.Println("address of intData: ", &intData)
	fmt.Println("what intPointer stores: ", intPointer)

	*intPointer = 30
	fmt.Println("what intData noew stores: ", intData)
}
