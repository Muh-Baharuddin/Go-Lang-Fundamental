package main

import "fmt"

func main() {
	languages := [...]string{
		"Ruby",
		"C++",
		"Typescript",
		"Java",
		"Go",
		"Javascript",
	}

	for index, length := range languages {
		fmt.Println("index:", index, "language:", length)
	}
}