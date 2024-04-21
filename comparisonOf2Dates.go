package main

import (
	"fmt"
	"time"
)

func main() {
	date1 := time.Date(2020, 2, 5, 5, 20, 0, time.UTC)
	date2 := time.Date(2020, 2, 5, 5, 20, 0, time.UTC)
	date3 := time.Date(2020, 2, 7, 5, 20, 0, time.UTC)

	if date1.Equal(date2) {
		fmt.Println("date1 and date2 are equal")
	} else {
		fmt.Println("date1 and date2 are not equal")
	}

	if date1.Equal(date3) {
		fmt.Println("date1 and date3 are equal")
	} else {
		fmt.Println("date1 and date3 are not equal")
	}

}
