package main

import (
	"fmt"
	"sort"
)

func main() {
	// numbers := []int{4, 2, 7, 1, 9, 3}

	// sort.Ints(numbers)
	// fmt.Println("Sorted: ", numbers)

	words := []string{"banana", "apple", "orange", "grapes", "pineapple"}

	sort.Strings(words)
	fmt.Println("Sorted:", words)
}
