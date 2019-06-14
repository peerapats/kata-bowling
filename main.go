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

	bonus string
	score3 int64

	totalScore int64
	currentScore int64
	sumScore int64
}

func main() {
	scores := [...] string{
		// "1", "6", 
		// "3", "/",
		// "X", "-",
		// "-", "/",
		// "1", "6",
		// "1", "6",
		// "1", "6",
		// "X", "-",
		// "X", "-",
		// "X", "-", "2",

		// "6", "1",
		// "9", "-",
		// "8", "/",
		// "5", "/",
		// "8", "-",
		// "6", "2",
		// "9", "/",
		// "7", "2",
		// "8", "/",
		// "9", "/", "7",
		// // ----------> sum 127

		// "X", "-",
		// "X", "-",
		// "7", "/",
		// "8", "/",
		// "X", "0",
		// "9", "/",
		// "X", "-",
		// "X", "-",
		// "X", "-",
		// "X", "7", "/",
		// ----------> sum 232

		// "5", "/",
		// "8", "/",
		// "9", "/",
		// "7", "/",
		// "8", "/",
		// "6", "/",
		// "9", "/",
		// "7", "/",
		// "6", "/",
		// "4", "5", "-",
		// // ----------> sum 163

			"X", "-",
			"X", "-",
			"X", "-",
			"X", "-",
			"X", "-",
			"X", "-",
			"X", "-",
			"X", "-",
			"X", "-",
			"X", "X", "X",
		// 	// "10", "0",
		// 	// "10", "0",
		// 	// "10", "0",
		// 	// "10", "0",
		// 	// "10", "0",
		// 	// "10", "0",
		// 	// "10", "0",
		// 	// "10", "0",
		// 	// "10", "0",
		// 	// "10", "10",
		// 	// "10", "10",
		// 	// ----------> sum 300
	}
	var result = board(scores);
	// fmt.Println(result[0].totalScore);

	for index, frame := range result {
		fmt.Println("Frame:", index + 1, "===>", frame.name1, frame.name2, frame.bonus, ":", frame.score3,  frame.totalScore, "," , frame.sumScore )
	}
}

func board(scores [21]string) []Frame {
	var frames = makeFrames(scores);
	var result = printScore(frames);
	return result;
}

func printScore(frames []Frame) []Frame {
	var sum int64;
	for index, frame := range frames {
		var next1 = index + 1;
		var next2 = index + 2;
		// fmt.Println(index, next1, next2, len(frames));
		var isFrame10 = index + 1 == 10;
		if(isFrame10){
			// spare - plus bonus 1 box
			if(frame.name2 == "/"){
				frame.totalScore += frames[index].score3;
			} 
			// spare - plus bonus 3 box
			if(frame.name1 == "X"){
				frame.totalScore += frames[index].score3;
			}
		} else {
			// spare - plus bonus 1 box
			if(frame.name2 == "/"){
				frame.totalScore += frames[next1].score1; // first box
			}
			// strike - plus bonus 2 box
			if(frame.name1 == "X"){
				// first box
				frame.totalScore += frames[next1].score1; 
				// second box
				var isFrame9 = index + 1 == 9;
				if(frames[next1].score2 == 0 &&  frames[next1].score1 + frames[next1].score2 == 10 && !isFrame9) {
					frame.totalScore += frames[next2].score1;
				} else {
					frame.totalScore += frames[next1].score2;
				}
			}
		}
 
		sum += frame.totalScore
		frame.sumScore += sum;
		frames[index] = frame;
	}
	return frames;
}

func makeFrames(scores [21]string) []Frame {
	var frames []Frame
	for index := range scores {
		var isEndFrame = index > 0 && index % 2 != 0;
		// fmt.Println("Index", index, "Value ", scores[index]);
		var isFrame10 = index == 19;
		var isBonus = index == 21;
		if(!isEndFrame && !isBonus){
			continue;
		}

		var symbol1 = scores[index - 1];
		var symbol2 = scores[index];
		var f = new(Frame);

		var score1 = convertBowling1(symbol1);
		var score2 = convertBowling2(score1, symbol2);
		var sumScore = score1 + score2;

		f.name1 = symbol1; 
		f.name2 = symbol2;
		f.score1 = score1;
		f.score2 = score2;
		f.totalScore = sumScore;
		f.currentScore = sumScore;

		if(isFrame10){
			var bonus = scores[index + 1];
			var score3 = convertBowling2(score2, bonus);
			f.bonus = bonus;
			f.score3 = score3;
		}
		frames = append(frames, *f);
	}
	
	// for index, frame := range frames {
	// 	fmt.Println("Frame:", index + 1, "===>",frame.name1, frame.name2, frame.score3,  ":", frame.score1, frame.score2,);
	// }
	return frames;
}

func convertBowling1(symbol1 string) int64 {
	if(symbol1 == "X"){
		return 10;
	}
	if(symbol1 == "-"){
		return 0;
	}
	
	score1, err := strconv.ParseInt(symbol1, 10, 64)
	if err != nil {
		fmt.Println(err)
		return 0;
	}	
	
	return score1;
}

func convertBowling2(score1 int64, symbol2 string) int64 {
	if(symbol2 == "/") {
		return 10 - score1;
	}

	var score2 = convertBowling1(symbol2);
	if(score1 == 10 && (score1 + score2 == 10) ){
		return 0;
	}
	return convertBowling1(symbol2);
}