package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var person1 Person
	var person2 Person

	person1.name = "Hege"
	person1.age = 45
	person1.job = "Teacher"
	person1.salary = 6000

	person2.name = "Cecilie"
	person2.age = 24
	person2.job = "Marketing"
	person2.salary = 4500

	fmt.Println("Name: ", person1.name)
	fmt.Println("Age: ", person1.age)
	fmt.Println("Job: ", person1.job)
	fmt.Println("Salary: ", person1.salary)
	fmt.Println("--------------------------")
	fmt.Println("Name: ", person2.name)
	fmt.Println("Age: ", person2.age)
	fmt.Println("Job: ", person2.job)
	fmt.Println("Salary: ", person2.salary)

}
