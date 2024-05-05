package main

import "fmt"

func main() {
	// personAge := map[string]int{"Hermione": 21, "Harry": 20, "John": 25} //duplicate key can not be used
	// fmt.Println("Initial map: ", personAge)

	// delete(personAge, "John")
	// fmt.Println("Updated map: ", personAge)

	//test that a key is present in map with a two-value assignment
	m := make(map[string]int)

	m["answer"] = 42
	fmt.Println("the value: ", m["answer"])

	m["answer"] = 48
	fmt.Println("the value: ", m["answer"])

	delete(m, "answer")
	fmt.Println("the value: ", m["answer"])

	v, ok := m["answer"]
	fmt.Println("the value: ", v, "present?", ok)
}
