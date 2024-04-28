package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "india is great country"
	var result1 string
	var result2 string

	result1 = strings.Replace(str, "i", "I", -1) //replace all occurence
	fmt.Println("result1: ", result1)

	result2 = strings.Replace(str, "i", "I", 2) //replace only 2 occurance
	fmt.Println("result2: ", result2)
}
