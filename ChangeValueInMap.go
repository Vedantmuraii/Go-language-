package main

import "fmt"

func main() {
	capital := map[string]string{"Nepal": "Kathmandu", "US": "New York"}
	fmt.Println("Initial map:", capital)

	capital["US"] = "Washington DC"
	fmt.Println("Updated Map: ", capital)
}
