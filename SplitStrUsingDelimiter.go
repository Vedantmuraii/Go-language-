package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "india-is-a-great-country"
	var strArr []string = strings.Split(str, "-")

	fmt.Println("spilliting string:")
	for i := 0; i < length(strArr); i++ {
		fmt.Println(strArrr[i])
	}
}
