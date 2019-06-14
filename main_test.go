
package main

import "testing"
import "github.com/stretchr/testify/assert"
import (
	"fmt"
)

type Person struct {
	name string
}

func TestFrame1DoubleStrike(t *testing.T) {
	scores := [...] string{
		"X", "-",
		"X", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-", "-", 
	}
	result := board(scores);
	expected := int64(30);
	printFrames(1, expected, result);
	assert.Equal(t, expected, result[10 - 1].sumScore, "Result should be 30")
}

func printFrames(frame int64, expected int64,  frames []Frame) {
	fmt.Printf("------------------- Test Frame: %d Expected: %d -------------------" , frame, expected);
	fmt.Println("");
	for index, frame := range frames {
		fmt.Println("Frame:", index + 1, "==>", frame.name1, frame.name2, frame.bonus,  ":", frame.totalScore );
	}
}

func Test2(t *testing.T) {
		scores := [...] string{
			"X", "-",
			"X", "-",
			"2", "-",
			"-", "-",
			"-", "-",
			"-", "-",
			"-", "-",
			"-", "-",
			"-", "-",
			"-", "-", "-",
		}
		result := board(scores);
		expected := int64(36);
		printFrames(1, expected, result);
		assert.Equal(t, expected, result[10 - 1].sumScore, "Result should be 1")
}

func Test3(t *testing.T) {
	scores := [...] string{
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"X", "-",
		"9", "/", "-",
	}
	result := board(scores);
	expected := int64(20);
	printFrames(9, expected, result);
	assert.Equal(t, expected, result[9 - 1].totalScore, "Result should be 1")
}

func Test4(t *testing.T) {
	scores := [...] string{
		"X", "-",
		"X", "-",
		"9", "/",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-", "-",
	}
	result := board(scores);
	var expected int64;
	expected = 29;
	assert.Equal(t, expected, result[1 - 1].totalScore, "Result should be 1")
}

func Test5(t *testing.T) {
	scores := [...] string{
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"X", "-",
		"X", "-", "-",
	}
	result := board(scores);
	var expected int64;
	expected = 20;
	assert.Equal(t, expected, result[9 - 1].totalScore, "Result should be 1")
}

func Test6(t *testing.T) {
	scores := [...] string{
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"X", "-",
		"X", "-", "7",
	}
	result := board(scores);
	var expected int64;
	expected = 17;
	assert.Equal(t, expected, result[10 - 1].totalScore, "Result should be 1")
}

func Test7(t *testing.T) {
	scores := [...] string{
		"X", "-",
		"X", "-",
		"X", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-", "-",
	}
	result := board(scores);
	var expected int64;
	expected = 30;
	assert.Equal(t, expected, result[1 - 1].totalScore, "Result should be 1")
}

func Test8(t *testing.T) {
	scores := [...] string{
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"X", "-",
		"X", "X", "X",
	}
	result := board(scores);
	var expected int64;
	expected = 30;
	assert.Equal(t, expected, result[10 - 1].totalScore, "Result should be 1")
}

func TestFrame10DoubleStrike(t *testing.T) {
	scores := [...] string{
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"X", "-",
		"X", "2", "3",
	}
	result := board(scores);
	var expected int64;
	expected = 37;
	assert.Equal(t, expected, result[10 - 1].sumScore, "Result should be 37")
}

func TestFrame10Spare(t *testing.T) {
	scores := [...] string{
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "/", "7",
	}
	result := board(scores);
	var expected int64;
	expected = 17;
	assert.Equal(t, expected, result[10 - 1].sumScore, "Result should be 17")
}

func TestFrame10Strike(t *testing.T) {
	scores := [...] string{
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"-", "-",
		"X", "X", "5",
	}
	result := board(scores);
	var expected int64;
	expected = 25;
	assert.Equal(t, expected, result[10 - 1].sumScore, "Result should be 25")
}