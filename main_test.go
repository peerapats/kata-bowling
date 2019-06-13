
package main

import "testing"
import "github.com/stretchr/testify/assert"
// import (
// 	"fmt"
// )

type Person struct {
	name string
}

func Test1(t *testing.T) {
	scores := [...] string{
		"10", "0",
		"7", "2",
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
	}
	result := board(scores);
	var expected int64;
	expected = 19;
	assert.Equal(t, expected, result[0].totalScore, "Result should be 1")
}

func Test2(t *testing.T) {
		scores := [...] string{
			"10", "0",
			"10", "0",
			"2", "0",
			"0", "0",
			"0", "0",
			"0", "0",
			"0", "0",
			"0", "0",
			"0", "0",
			"0", "0",
			"0", "0",
		}
		result := board(scores);
		var expected int64;
		expected = 22;
		assert.Equal(t, expected, result[0].totalScore, "Result should be 1")
}

func Test3(t *testing.T) {
	scores := [...] string{
		"10", "0",
		"10", "0",
		"2", "0",
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
		"10", "0",
		"9", "1",
		"0", "0",
	}
	result := board(scores);
	var expected int64;
	expected = 20;
	assert.Equal(t, expected, result[8].totalScore, "Result should be 1")
}

func Test4(t *testing.T) {
	scores := [...] string{
		"10", "0",
		"10", "0",
		"9", "1",
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
	}
	result := board(scores);
	var expected int64;
	expected = 29;
	assert.Equal(t, expected, result[0].totalScore, "Result should be 1")
}

func Test5(t *testing.T) {
	scores := [...] string{
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
		"10", "0",
		"10", "0",
		"0", "0",
	}
	result := board(scores);
	var expected int64;
	expected = 20;
	assert.Equal(t, expected, result[8].totalScore, "Result should be 1")
}

func Test6(t *testing.T) {
	scores := [...] string{
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
		"10", "0",
		"10", "0",
		"7", "2",
	}
	result := board(scores);
	var expected int64;
	expected = 17;
	assert.Equal(t, expected, result[9].totalScore, "Result should be 1")
}


func Test7(t *testing.T) {
	scores := [...] string{
		"10", "0",
		"10", "0",
		"10", "0",
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
		"0", "0",
	}
	result := board(scores);
	var expected int64;
	expected = 30;
	assert.Equal(t, expected, result[0].totalScore, "Result should be 1")
}
