package main

import (
	"fmt"
	"strconv"
	"strings"
)

type frame struct {
	pins_1   string
	pins_2   string
	pins_3   string
	score    int
	isStrike bool
	isSpare  bool
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func getScoreFrames(rolls []string) (frames []frame) {
	size := len(rolls)

	for i := 0; i < size; i++ {
		pins := strings.Split(rolls[i], "")

		var f = new(frame)
		f.pins_1 = pins[0]

		if f.pins_1 == "x" {
			f.score = 10
			f.isStrike = true
		} else if isNumeric(f.pins_1) {
			p1_score, _ := strconv.Atoi(f.pins_1)
			f.score = p1_score
		}

		if len(pins) > 1 {
			f.pins_2 = pins[1]
			if f.pins_2 == "/" {
				f.score = 10
				f.isSpare = true
			} else if isNumeric(f.pins_2) {
				p2_score, _ := strconv.Atoi(f.pins_2)
				f.score = f.score + p2_score
			}
		}

		frames = append(frames, *f)

		if i != 0 && frames[i-1].isStrike {
			frames[i-1].score = frames[i-1].score + f.score
		}
	}

	return frames
}

func main() {
	rolls := []string{"24", "x", "3/", "12", "x", "5/", "x", "x", "44", "xxx"}
	frames := getScoreFrames(rolls)
	for i := 0; i < len(frames); i++ {
		f := frames[i]
		fmt.Println("Frame", i+1, "["+f.pins_1+"]["+f.pins_2+"]", " = ", f.score)
	}
}
