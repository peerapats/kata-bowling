
package main

import "testing"
import "github.com/stretchr/testify/assert"

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
	assert.Equal(t, expected, result[10 - 1].sumScore, "Result should be 30")
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
	expected := int64(30);
	assert.Equal(t, expected, result[10 - 1].sumScore, "Result should be 1")
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
	expected := int64(59);
	assert.Equal(t, expected, result[10 - 1].sumScore, "Result should be 1")
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
	expected := int64(30);
	assert.Equal(t, expected, result[10 - 1].sumScore, "Result should be 1")
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
	expected := int64(37);
	assert.Equal(t, expected, result[10 - 1].sumScore, "Result should be 1")
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
	expected := int64(60);
	assert.Equal(t, expected, result[10 - 1].sumScore, "Result should be 1")
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
	expected := int64(60);
	assert.Equal(t, expected, result[10 - 1].sumScore, "Result should be 1")
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
	expected := int64(37);
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
	expected := int64(17);
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

func TestCovertErrorReturnPointZero(t *testing.T) {
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
		"*", "X", "5",
	}
	result := board(scores);
	var expected int64;
	expected = 10;
	assert.Equal(t, expected, result[10 - 1].sumScore, "Result should be 10")
}