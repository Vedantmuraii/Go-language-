package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "India"
	var substr string = "ia"
	var result bool = true

	result = strings.HasSuffix(str, substr)

	if result == true {
		fmt.Println(str, "has suffix", substr)
	} else {
		fmt.Println(str, "does not have", substr, "as suffix")
	}
}
