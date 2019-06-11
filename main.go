package main

import (
	"fmt"
	"strconv"
)

// func cal(val string, next string) {
// 	points := strings.Split(val, "")
// 	frame := 0
// 	for key, value := range points {
// 		if ()
// 	}
// }

func main() {
	var games = [...]string{"24", "02", "23", "15", "13", "10", "53", "63", "27", "11"}
	fmt.Println("Input => ", games)
	totalScore := int64(0)
	for key, value := range games {
		i1, _ := strconv.ParseInt(string(value[0]), 10, 64)
		i2, _ := strconv.ParseInt(string(value[1]), 10, 64)
		frameScore := int64(0)
		frameScore = i1 + i2
		totalScore += frameScore
		fmt.Println("Frame", key+1, " = ", i1+i2)
	}
	fmt.Println(totalScore)
}
