package main

import "fmt"

func main() {
	var gamingConsole []string

	gamingConsole = append(gamingConsole, "Playstation")
	gamingConsole = append(gamingConsole, "Nintendo Switch")
	gamingConsole = append(gamingConsole, "Xbox One")

	for _, console := range gamingConsole {
		fmt.Println(console)
	}

}