package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5) //make = memory allocate krta hai, 3 is the length and 5 is the capacity, len<=cap
	printSlice(numbers)
}
func printSlice(x []int) {
	fmt.Printf("Len = %d cap = %d slice = %v\n", len(x), cap(x), x)
}
