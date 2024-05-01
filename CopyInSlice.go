package main

import "fmt"

func main() {
	primeNumbers := []int{2, 3, 5, 7}
	numbers := []int{1, 2, 3, 4, 5, 6}

	copy(numbers, primeNumbers) //target and then source
	//prime numbers ke 3 numbers copy/replace hue hai numbers mei kyuki numbers ki capacity 3 thi
	fmt.Println("Numbers: ", numbers)
}
