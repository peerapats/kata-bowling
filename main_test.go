
package main

import "testing"
import "github.com/stretchr/testify/assert"

type Person struct {
	name string
}

func Test1(t *testing.T) {
// Frame: 1 ===> X - ( 10 ) = 19 , 19
// Frame: 2 ===> 7 2 ( 9 ) = 9 , 28
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
	// Frame: 1 ===> X - ( 10 ) = 19 , 19
	// Frame: 2 ===> 7 2 ( 9 ) = 9 , 28
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