package main

import (
	"fmt"
	"strconv"
)

// เก็บคะแนนของแต่ละเกมส์
var score [10]int

// เก็บตัวทดการสะสมคะแนน
var carry [10]int

var input = [10]string{
	"X-",
	"X-",
	"X-",
	"X-",
	"X-",
	"X-",
	"X-",
	"X-",
	"X-",
	"XXX",
}

func main() {

	fmt.Println("Input: ", input)

	for gameNoIndex, value := range input {

		for _, ch := range value {
			str := string(ch)

			var scoreValue int
			if str == "X" || str == "/" {
				scoreValue = 10
			} else if str == "-" {
				scoreValue = 0
			} else {
				iValue, _ := strconv.ParseInt(str, 10, 32)
				scoreValue = int(iValue)
			}

			// Fullfill score to preivos Strike or Spare game
			fullfillScore(scoreValue)

			// Keep current score value.
			score[gameNoIndex] += scoreValue

			// Keep carry fullfill score to current Game
			if gameNoIndex < len(input)-1 {
				if str == "X" {
					carry[gameNoIndex] = 2
					break
				} else if str == "/" {
					carry[gameNoIndex] = 1
				}
			}
		}

	}

	showScore()
	fmt.Println("Total Score: ", sumScore())
}

func fullfillScore(scoreValue int) {
	for index, value := range carry {
		if value > 0 {
			score[index] += scoreValue
			carry[index]--
		}
	}
}

func showScore() {
	for index, value := range score {
		fmt.Printf("Frame %d: %d\n", index+1, value)
		// fmt.Println("Frame %d: %d", index+1, value)
	}
}

func sumScore() int {
	totalScore := 0
	for _, val := range score {
		totalScore += val
	}

	return totalScore
}
