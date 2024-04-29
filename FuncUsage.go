// to illustrate the usage of 2 var in one line
/*package main

import (
	"fmt"
)

func main() {
	var i, j = 101, "hello"
	fmt.Print(i, " ", j)
}*/

// to liiustrate the usage of fmt.Println() func
package main

import (
	"fmt"
)

func main() {
	const name, dept = "GeeksforGeeks", "CS"
	fmt.Println(name, "is", "a", dept, "Portal.")
}
