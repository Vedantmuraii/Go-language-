package main

import (
	"fmt"
	"strings"
)

func main() {
	str := []string{"India", "is", "a", "great", "country"}
	var result string
	result = strings.Join(str, "_")
	fmt.Println("result: ", result)
}
