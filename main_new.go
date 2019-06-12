package main_new

import (
	"fmt"
	"strconv"
)

// เก็บคะแนนของแต่ละเกมส์
// var score [10]int

// // เก็บตัวทดการสะสมคะแนน
// var carry [10]int



func main() {

	

	var input = [10]string{
		"--",  
		"--",  
		"--",  
		"--",  
		"--",
		"--",
		"--",
		"--",
		"--",
		"5/9",
	};

	fmt.Println("Input: ", input)

	var result = playGame(input);

	fmt.Println(result);

	//showScore()
	fmt.Println("Total Score: ", sumScore(result));
}

func playGame(input [10]string) [10]int {
	// เก็บคะแนนของแต่ละเกมส์
	var score [10]int

	// เก็บตัวทดการสะสมคะแนน
	var carry [10]int

	for gameNoIndex, value := range input {

		for _, ch := range value {
			str := string(ch)

			var scoreValue int
			if str == "-" {
				scoreValue = 0
			} else if str == "X" || str == "/" {
				scoreValue = 10

				if gameNoIndex == 9 {
					score[gameNoIndex] += scoreValue
				} else {
					score[gameNoIndex] = scoreValue
				}
			} else {
				iValue, _ := strconv.ParseInt(str, 10, 32)
				scoreValue = int(iValue)

				score[gameNoIndex] += scoreValue
			}

			// Fullfill score to preivos Strike or Spare game
			fullfillScore(scoreValue, &score, &carry);

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

	return score;
}

func fullfillScore(scoreValue int, score *[10]int, carry *[10]int) {
	for index, value := range carry {
		if value > 0 {
			score[index] += scoreValue
			carry[index]--
		}
	}
}

// func showScore() {
// 	for index, value := range score {
// 		fmt.Printf("Frame %d [ %s ]: %d\n", index+1, input[index], value)
// 		// fmt.Println("Frame %d: %d", index+1, value)
// 	}
// }
func sumScore(score [10]int) int {
	totalScore := 0
	for _, val := range score {
		totalScore += val
	}

	return totalScore
}

