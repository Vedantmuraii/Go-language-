package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Books
	var Book2 Books
	Book1.title = "Go Programming"
	Book1.author = "Mahesh Kumar"
	Book1.subject = "Go Programming Tutorials"
	Book1.book_id = 534879

	Book2.title = "Telecom Biling"
	Book2.author = "Zara Ali"
	Book2.subject = "Telecom Billing Tutorials"
	Book2.book_id = 9862487

	fmt.Printf("Book1 Title: %s\n", Book1.title)
	fmt.Printf("Book1 Author: %s\n", Book1.author)
	fmt.Printf("Book1 Subject: %s\n", Book1.subject)
	fmt.Printf("Book1 Book_id: %d\n", Book1.book_id)

	fmt.Printf("Book2 Title: %s\n", Book2.title)
	fmt.Printf("Book2 Author: %s\n", Book2.author)
	fmt.Printf("Book2 Subject: %s\n", Book2.subject)
	fmt.Printf("Book2 Book_id: %d\n", Book2.book_id)

}
