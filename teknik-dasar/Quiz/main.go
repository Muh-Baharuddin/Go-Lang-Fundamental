package main

import (
	"fmt"
	"quiz/multiply"
)

func main() {
	fmt.Println("Ini Quiz")
	perkalian := multiply.Multiply(5, 5)
	fmt.Println(perkalian)
}