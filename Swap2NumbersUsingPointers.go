package main

import "fmt"

func swap(x, y *int) {
	t := 0
	t = *x
	*x = *y
	*y = t
}
func main() {
	a, b := 10, 20
	fmt.Println("before swap: ", a, b)
	swap(&a, &b)
	fmt.Println("after swap: ", a, b)
}
