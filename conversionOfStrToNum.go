package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	var num int64

	fmt.Print("Enter string num: ")
	fmt.Scanf("%s", &str)

	num, _ = strconv.ParseInt(str, 0, 64) //num, _ = num, error('_' error ke liye hai but kyuki error nhi chahiye toh _ likh dete hai)
	fmt.Println("Num: ", num)             //3 arguments(str,0,64) i.e. str jo likh rhe hai, 0 = base, and 64 is the size
}
