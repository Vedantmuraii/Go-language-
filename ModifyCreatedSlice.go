package main

import "fmt"

func main() {
	arr := [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	intSlice := arr[2:5]
	for _, ele := range intSlice {
		fmt.Printf("%d\n", ele)
	}
	fmt.Println("-------------")
	intSlice[1] = 500
	fmt.Println("Slice elements: ")
	for _, ele := range intSlice {
		fmt.Printf("%d\n", ele)
	}
}
