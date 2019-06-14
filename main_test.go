package main

import "testing"
import "github.com/stretchr/testify/assert"
import "fmt"

func TestAlwaysReturn52WithAssert(t *testing.T) {

	score := []string{"24", "44", "41", "26", "11", "13", "31", "41", "26", "11"}
	scoreBoard := getScoreBoard(score)
	// frames := scoreBoard.frames

	for i := 0; i < len(scoreBoard.frames); i++ {
		f := scoreBoard.frames[i]
		if i == 9 {
			fmt.Println("Frame", i+1, "["+f.pins_1+"]["+f.pins_2+"]["+f.pins_3+"]", "=", f.score)
		} else {
			fmt.Println("Frame", i+1, "["+f.pins_1+"]["+f.pins_2+"]", "=", f.score)
		}
	}

	fmt.Println("Total =", scoreBoard.totalScore)
	expected := []Frame{
		Frame{"2", "4", "", 6, false, false},
		Frame{"4", "4", "", 8, false, false},
		Frame{"4", "1", "", 5, false, false},
		Frame{"2", "6", "", 8, false, false},
		Frame{"1", "1", "", 2, false, false},
		Frame{"1", "3", "", 4, false, false},
		Frame{"3", "1", "", 4, false, false},
		Frame{"4", "1", "", 5, false, false},
		Frame{"2", "6", "", 8, false, false},
		Frame{"1", "1", "", 2, false, false},
	}
	assert.Equal(t, expected, scoreBoard.frames, "total score should equal 52")
	fmt.Println("total score should equal 52")
}

func TestAlwaysReturn0WithAssert(t *testing.T) {

	score := []string{"--", "--", "--", "--", "--", "--", "--", "--", "--", "--"}
	scoreBoard := getScoreBoard(score)
	// frames := scoreBoard.frames

	for i := 0; i < len(scoreBoard.frames); i++ {
		f := scoreBoard.frames[i]
		if i == 9 {
			fmt.Println("Frame", i+1, "["+f.pins_1+"]["+f.pins_2+"]["+f.pins_3+"]", "=", f.score)
		} else {
			fmt.Println("Frame", i+1, "["+f.pins_1+"]["+f.pins_2+"]", "=", f.score)
		}
	}

	fmt.Println("Total =", scoreBoard.totalScore)
	expected := []Frame{
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
	}
	assert.Equal(t, expected, scoreBoard.frames, "total score should equal 0")
	fmt.Println("total score should equal 0")
}

func TestAlwaysIsSpareReturn14WithAssert(t *testing.T) {

	score := []string{"4/", "--", "--", "--", "--", "--", "--", "--", "--", "--"}
	scoreBoard := getScoreBoard(score)
	for i := 0; i < len(scoreBoard.frames); i++ {
		f := scoreBoard.frames[i]
		if i == 9 {
			fmt.Println("Frame", i+1, "["+f.pins_1+"]["+f.pins_2+"]["+f.pins_3+"]", "=", f.score)
		} else {
			fmt.Println("Frame", i+1, "["+f.pins_1+"]["+f.pins_2+"]", "=", f.score)
		}
	}

	fmt.Println("Total=", scoreBoard.totalScore)
	expected := []Frame{
		Frame{"4", "/", "", 10, false, true},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
	}
	assert.Equal(t, expected, scoreBoard.frames, "total score should equal 10")
	fmt.Println("total score should equal 10")
}

func TestAlwaysIsStrikeReturnWithAssert(t *testing.T) {

	score := []string{"x", "4/", "--", "--", "--", "--", "--", "--", "--", "--"}
	scoreBoard := getScoreBoard(score)
	for i := 0; i < len(scoreBoard.frames); i++ {
		f := scoreBoard.frames[i]
		if i == 9 {
			fmt.Println("Frame", i+1, "["+f.pins_1+"]["+f.pins_2+"]["+f.pins_3+"]", "=", f.score)
		} else {
			fmt.Println("Frame", i+1, "["+f.pins_1+"]["+f.pins_2+"]", "=", f.score)
		}
	}

	fmt.Println("Total =", scoreBoard.totalScore)
	expected := []Frame{
		Frame{"x", "", "", 20, true, false},
		Frame{"4", "/", "", 10, false, true},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
	}
	assert.Equal(t, expected, scoreBoard.frames, "total score should equal 30")
	fmt.Println("total score should equal 30")
}

func TestAlwaysReturnWithAssert(t *testing.T) {

	score := []string{"--", "--", "--", "--", "--", "--", "--", "--", "--", "x"}
	scoreBoard := getScoreBoard(score)
	for i := 0; i < len(scoreBoard.frames); i++ {
		f := scoreBoard.frames[i]
		if i == 9 {
			fmt.Println("Frame", i+1, "["+f.pins_1+"]["+f.pins_2+"]["+f.pins_3+"]", "=", f.score)
		} else {
			fmt.Println("Frame", i+1, "["+f.pins_1+"]["+f.pins_2+"]", "=", f.score)
		}
	}

	fmt.Println("Total =", scoreBoard.totalScore)
	expected := []Frame{
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"x", "", "", 10, true, false},
	}
	assert.Equal(t, expected, scoreBoard.frames, "total score should equal 10")
	fmt.Println("total score should equal 10")
}

func TestAlwaysReturnFrame10AStrikeWithAssert(t *testing.T) {

	score := []string{"--", "--", "--", "--", "--", "--", "--", "--", "--", "x-5"}
	scoreBoard := getScoreBoard(score)
	for i := 0; i < len(scoreBoard.frames); i++ {
		f := scoreBoard.frames[i]
		if i == 9 {
			fmt.Println("Frame", i+1, "["+f.pins_1+"]["+f.pins_2+"]["+f.pins_3+"]", "=", f.score)
		} else {
			fmt.Println("Frame", i+1, "["+f.pins_1+"]["+f.pins_2+"]", "=", f.score)
		}
	}

	fmt.Println("Total =", scoreBoard.totalScore)
	expected := []Frame{
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"x", "-", "5", 15, true, false},
	}
	assert.Equal(t, expected, scoreBoard.frames, "total score should equal 12")
	fmt.Println("total score should equal 15")
}

func TestAlwaysReturnFrame10DoubleStrikeWithAssert(t *testing.T) {

	score := []string{"--", "--", "--", "--", "--", "--", "--", "--", "--", "x-x"}
	scoreBoard := getScoreBoard(score)
	for i := 0; i < len(scoreBoard.frames); i++ {
		f := scoreBoard.frames[i]
		if i == 9 {
			fmt.Println("Frame", i+1, "["+f.pins_1+"]["+f.pins_2+"]["+f.pins_3+"]", "=", f.score)
		} else {
			fmt.Println("Frame", i+1, "["+f.pins_1+"]["+f.pins_2+"]", "=", f.score)
		}
	}

	fmt.Println("Total =", scoreBoard.totalScore)
	expected := []Frame{
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"-", "-", "", 0, false, false},
		Frame{"x", "-", "x", 10, true, false},
	}
	assert.Equal(t, expected, scoreBoard.frames, "total score should equal 10")
	fmt.Println("total score should equal 10")
}
