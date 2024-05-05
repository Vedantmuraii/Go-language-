package main

import "fmt"

func main() {
	// squaredNumber := map[int]int{2: 4, 3: 9, 4: 16, 5: 25}
	// for number, squared := range squaredNumber {
	// 	fmt.Printf("square of %d is %d\n", number, squared)
	// }

	places := map[string]string{"nepal": "kathmandu", "us": "washington dc", "norway": "oslo"}
	for _, capital := range places {
		fmt.Println(capital)
	}

}
