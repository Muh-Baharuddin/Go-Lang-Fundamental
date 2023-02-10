package main

import "fmt"

func main() {
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	//quiz 1
	var total int
	for _, score := range scores {
		total += score
	}
	length := len(scores)
	average := float32(total)/float32(length)
	fmt.Println(average)

	//quiz 2
	var goodScores []int

	for _, score := range scores {
		if score >= 90 {
			goodScores = append(goodScores, score)
		}
	}

	fmt.Println(goodScores)
}