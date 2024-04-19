package main

import "fmt"

func main() {
	students := map[int]string{1: "John", 2: "Lily"}
	fmt.Println("Initial map: ", students)

	students[3] = "Robin"
	students[5] = "Julie"
	fmt.Println("Updated map: ", students)
}
