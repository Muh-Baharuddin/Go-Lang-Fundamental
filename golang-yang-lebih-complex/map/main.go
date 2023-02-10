package main

import "fmt"

func main() {
	myMap := map[string]int{
		"Typescript": 100,
		"Java":       80,
		"Go":         90,
	}

	value1, isAvailable1 := myMap["Python"]
	value2, isAvailable2 := myMap["Go"]
	fmt.Println(isAvailable1)
	fmt.Println(value1)
	fmt.Println(isAvailable2)
	fmt.Println(value2)
}