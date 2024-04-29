package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "ABC PQR LMN PQR"
	str2 := "XYZ LMN LMN PQR"

	fmt.Println(strings.ReplaceAll(str1, "PQR", "TUV"))
	fmt.Println(strings.ReplaceAll(str2, "XYZ", "ABC"))
}
