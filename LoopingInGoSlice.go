package main

import "fmt"

func main() {
	// numbers := []int{2, 4, 6, 8, 10}
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }
	//addition in slice

	evenNumbers := []int{2, 4}
	oddNumbers := []int{1, 3}
	evenNumbers = append(evenNumbers, oddNumbers...)
	for i := 0; i < len(evenNumbers); i++ {
		fmt.Println(evenNumbers[i])
	}
}
