package main

import (
	"strconv"
	"strings"
)

type Frame struct {
	pins_1   string
	pins_2   string
	pins_3   string
	score    int
	isStrike bool
	isSpare  bool
}

type ScoreBoard struct {
	totalScore int
	frames     []Frame
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func toNumber(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

func sumScores(frames []Frame) (sum int) {
	sum = 0
	for i := 0; i < len(frames); i++ {
		sum += frames[i].score
	}
	return sum
}

func getScoreBoard(rolls []string) ScoreBoard {
	var frames []Frame
	for i := 0; i < len(rolls); i++ {
		pins := strings.Split(rolls[i], "")

		var f = new(Frame)
		f.pins_1 = pins[0]

		if f.pins_1 == "x" {
			f.score = 10
			f.isStrike = true
		} else if isNumeric(f.pins_1) {
			f.score = toNumber(f.pins_1)
		}

		if len(pins) > 1 {
			f.pins_2 = pins[1]
			if f.pins_2 == "/" {
				f.score = 10
				f.isSpare = true
			} else if isNumeric(f.pins_2) {
				p2_score := toNumber(f.pins_2)
				f.score = f.score + p2_score
			}
		}

		// frame 10
		if i == 9 && (f.isStrike || f.isSpare) && len(pins) > 2 {
			f.pins_3 = pins[2]
			if f.pins_3 == "x" || f.pins_3 == "/" {
				f.score = 10
			} else {
				p3_score := toNumber(f.pins_3)
				f.score = f.score + p3_score
			}
		}

		frames = append(frames, *f)

		if i != 0 && frames[i-1].isStrike {
			frames[i-1].score = frames[i-1].score + f.score
		}
	}

	totalScore := sumScores(frames)
	scoreBoard := ScoreBoard{totalScore: totalScore, frames: frames}
	return scoreBoard
}

func main() {
	// rolls := []string{"24", "x", "3/", "12", "x", "5/", "x", "x", "44", "xxx"}
	// scoreBoard := getScoreBoard(rolls)
	// for i := 0; i < len(scoreBoard.frames); i++ {
	// 	f := scoreBoard.frames[i]
	// 	if i == 9 {
	// 		fmt.Println("Frame", i+1, "["+f.pins_1+"]["+f.pins_2+"]["+f.pins_3+"]", "=", f.score)
	// 	} else {
	// 		fmt.Println("Frame", i+1, "["+f.pins_1+"]["+f.pins_2+"]", "=", f.score)
	// 	}
	// }
	// fmt.Println("Total score =", scoreBoard.totalScore)
}
