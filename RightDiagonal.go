package main

import "fmt"

func main() {
	var matrix [3][3]int
	fmt.Printf("enter matrix elements: \n")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("elements: matrix[%d][%d]: ", i, j)
			fmt.Scanf("%d", &matrix[i][j])
		}
	}
	fmt.Printf("Matrix: \n")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d", matrix[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("right diagonal of matrix: \n")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i+j == 2 {
				fmt.Printf("%d", matrix[i][j])
			} else {
				fmt.Printf(" ")
			}

		}
		fmt.Printf("\n")
	}

}
