package main

import "fmt"

func main() {
	students := []map[string]string{ //membuat slice yang isinya map
		{"name": "Aldi", "score": "A"},
		{"name": "Zaki", "score": "B"},
		{"name": "Alfian", "score": "E"},
	}

	fmt.Println(students)

	for _, student := range students {
		fmt.Println(student)
		fmt.Println(student["name"])
	}
}