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
	"-/",
	"3/",
	"X-",
	"XX8",
}

// มี Bug ใน Frame ที่ 8 จะต้องมีแต้มเป็น 20
// เพราะว่า Frame 9 เกิด Spare

func main() {

	fmt.Println("Input: ", input)

	for gameNoIndex, value := range input {

		for _, ch := range value {
			str := string(ch)

			var scoreValue int
			if str == "-" {
				scoreValue = 0
			} else if str == "X" {
				scoreValue = 10
			} else if str == "/" {
				var priorScore = 20
				if score[gameNoIndex] < 10 {
					priorScore = 10
				}

				scoreValue = priorScore - score[gameNoIndex]

			} else {
				iValue, _ := strconv.ParseInt(str, 10, 32)
				scoreValue = int(iValue)
			}

			// Fullfill score to preivos Strike or Spare game
			fullfillScore(scoreValue)

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
		fmt.Printf("Frame %d [ %s ]: %d\n", index+1, input[index], value)
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
