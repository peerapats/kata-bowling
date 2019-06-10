package main

import (
	"fmt"
)

func main() {
	var games = []string{"2/", "x", "x", "x", "x", "x", "x", "x", "x", "xxx"}
	score := 0
	for index, value := range games {
		fmt.Println(index, " = ", value)
	}
	fmt.Println(score)
}
