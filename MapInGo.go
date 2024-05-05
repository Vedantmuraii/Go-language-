package main

import "fmt"

func main() {

	// subjectMarks := map[string]float32{"Golang": 85, "Java": 80, "Python": 81}
	// fmt.Println(subjectMarks)
	// fmt.Println(subjectMarks["Golang"])

	// flowerColor := map[string]string{"sunflower": "yellow", "jasmine": "white", "hibiscus": "red"}
	// fmt.Println(flowerColor["sunflower"])
	// fmt.Println(flowerColor["hibiscus"])

	//make function
	// students := make(map[int]string)
	// students[1] = "Harry"
	// students[2] = "Lily"
	// students[3] = "Hermonie"

	// fmt.Println(students)
	// fmt.Printf("%T\n", students)

	//access keys in map
	places := map[string]string{"nepal": "kathmandu", "us": "washington dc", "norway": "oslo"}

	for country := range places {
		fmt.Println(country)
	}
}
