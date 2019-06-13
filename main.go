package main

import (
	"fmt"
	"strconv"
)

type Frame struct {
	name1 string
	name2 string
	score1 int64
	score2 int64
	totalScore int64
	currentScore int64
	sumScore int64
}

// http://www.star-circuit.com/article/another-article/bowling/Counting-bowling-score.html
func main() {
	scores := [...] string{
		"10", "0",
		"7", "2",
		"9", "1",
		"8", "1",
		"10", "0",
		"10", "0",
		"9", "1",
		"0", "10",
		"7", "2",
		"1", "10",
		"1", "0",
		// // ----------> sum 127

		// "6", "1",
		// "9", "0",
		// "8", "2",
		// "5", "5",
		// "8", "0",
		// "6", "2",
		// "9", "1",
		// "7", "2",
		// "8", "2",
		// "9", "1",
		// "7", "3",
		// // // ----------> sum 127

		// "10", "0",
		// "10", "0",
		// "7", "3",
		// "8", "2",
		// "10", "0",
		// "9", "1",
		// "10", "0",
		// "10", "0",
		// "10", "0",
		// "10", "7",
		// "3", "0",
		// // ----------> sum 232

		// "5", "5",
		// "8", "2",
		// "9", "1",
		// "7", "3",
		// "8", "2",
		// "6", "4",
		// "9", "1",
		// "7", "3",
		// "6", "4",
		// "4", "5",
		// "3", "0",
		// ----------> sum 163


		// "10", "0",
		// "10", "0",
		// "10", "0",
		// "10", "0",
		// "10", "0",
		// "10", "0",
		// "10", "0",
		// "10", "0",
		// "10", "0",
		// "10", "10",
		// "10", "10",
		// ----------> sum 300

	}

	var frames = makeFrames(scores);
	var result = printScore(frames);
	fmt.Println("==========>");
	for index, frame := range result {
		if(index + 1 > 10) {
			continue;
		}
		fmt.Println("Frame:", index + 1, "===>", frame.name1, frame.name2, "=", frame.totalScore, ",", frame.sumScore )
	}
}

func printScore(frames []Frame) []Frame {
	var sum int64;
	for index, frame := range frames {
		var next1 = index + 1;
		var next2 = index + 2;
		if(next1 > len(frames) || next2 > len(frames)){
			continue;
		}

		if(frame.name2 == "/"){
			frame.totalScore += frames[next1].score1;
		}

		if(frame.name1 == "X"){
			// find first bonus
			frame.totalScore += frames[next1].score1;
			var isFrame10 = index + 1 == 10;
			if(!isFrame10){
			// find second bonus
				if(frames[next1].score2 == 0){
					frame.totalScore += frames[next2].score1;
				} else {
					frame.totalScore += frames[next1].score2;
				}
			}
		}

		sum += frame.totalScore
		frame.sumScore += sum;

		frames[index] = frame;
		fmt.Println(index, next1, next2, len(frames));
	}
	return frames;
}

func makeFrames(scores [22]string) []Frame {
	var frames []Frame
	for index := range scores {
		var isEndFrame = index > 0 && index % 2 != 0;
		if(!isEndFrame){
			continue;
		}

		var score1 = scores[index - 1];
		var score2 = scores[index];
		var f = new(Frame);
		f.name1 = convertBowling1(score1);
		f.name2 = convertBowling2(score1, score2);
		f.score1 = scoreToInt(score1);
		f.score2 = scoreToInt(score2);
		f.totalScore = sumScore(score1, score2);
		f.currentScore = sumScore(score1, score2);
		frames = append(frames, *f);
	}
	return frames;
}

func convertBowling1(score1 string) string {
	if(score1 == "10"){
		return "X";
	}
	if(score1 == "0"){
		return "-";
	}
	return score1;
}

func convertBowling2(score1 string, score2 string) string {
	if(score1 == "10" && sumScore(score1, score2) == 10){
		return "-";
	}
	if(sumScore(score1, score2) == 10){
		return "/";
	}
	if(score2 == "0"){
		return "-";
	}
	if(score2 == "10"){
		return "X";
	}
	return score2;
}

func sumScore(score1 string, score2 string) int64 {
	return scoreToInt(score1) + scoreToInt(score2);
}

func scoreToInt(score string) int64 {
	num, err := strconv.ParseInt(score, 10, 0)
	if err != nil {
		fmt.Println(err)
		return 0;
	}	
	return num;
}