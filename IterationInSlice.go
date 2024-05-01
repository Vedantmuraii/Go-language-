package main

import "fmt"

func main() {
	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	intSlice := arr[2:5]
	fmt.Println(("Slice elements: "))
	// for index, ele := range intSlice {
	// 	fmt.Printf("index: %d, elements: %d\n", index, ele)
	// }
	for _, ele := range intSlice {
		fmt.Printf("elements: %d\n", ele)
	}
	fmt.Println("---------------")
	for index, _ := range intSlice {
		fmt.Printf("index: %d\n", index)
	}
	fmt.Println("--------------")
	for index, _ := range intSlice {
		fmt.Println(intSlice[index])
	}
}
