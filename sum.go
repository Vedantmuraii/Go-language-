package main

import (
	"fmt"
)

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
func main() {
	sum(10, 20)
	sum(10, 20, 30)
	sum(10, 20, 30, 40)
	nums := []int{10, 20, 30, 40}
	sum(nums...)
	sum()
}
