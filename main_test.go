package main

import "testing"
import "github.com/stretchr/testify/assert"
import "fmt"

func TestCaseZero(t *testing.T) {
	var score1 = [10]string{"--", "--", "--", "--", "--", "--", "--", "--", "--", "--"}
	var expect = [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	result := calc(score1)
	fmt.Println("")
	assert.Equal(t, expect, result, "Total should be 0")
}

func TestCaseSuperStrike(t *testing.T) {
	var score1 = [10]string{"x-", "x-", "x-", "x-", "x-", "x-", "x-", "x-", "x-", "xxx"}
	var expect = [10]int{30, 30, 30, 30, 30, 30, 30, 30, 30, 30}
	result := calc(score1)
	fmt.Println("")
	assert.Equal(t, expect, result, "Total should be 300")
}

func TestCaseNumber(t *testing.T) {
	var score1 = [10]string{"62", "--", "--", "--", "--", "--", "--", "--", "--", "--"}
	var expect = [10]int{8, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	result := calc(score1)
	fmt.Println("")
	assert.Equal(t, expect, result, "Total should be 8")
}

func TestCaseSpare(t *testing.T) {
	var score1 = [10]string{"4/", "52", "--", "--", "--", "--", "--", "--", "--", "--"}
	var expect = [10]int{15, 7, 0, 0, 0, 0, 0, 0, 0, 0}
	result := calc(score1)
	fmt.Println("")
	assert.Equal(t, expect, result, "Total should be 22")
}

func TestCaseStrike(t *testing.T) {
	var score1 = [10]string{"x-", "11", "--", "--", "--", "--", "--", "--", "--", "--"}
	var expect = [10]int{12, 2, 0, 0, 0, 0, 0, 0, 0, 0}
	result := calc(score1)
	fmt.Println("")
	assert.Equal(t, expect, result, "Total should be 14")
}

func TestCaseDoubleStrike(t *testing.T) {
	var score1 = [10]string{"x-", "x-", "5-", "--", "--", "--", "--", "--", "--", "--"}
	var expect = [10]int{25, 15, 5, 0, 0, 0, 0, 0, 0, 0}
	result := calc(score1)
	fmt.Println("")
	assert.Equal(t, expect, result, "Total should be 45")
}

func TestCaseLastTripleStrike(t *testing.T) {
	var score1 = [10]string{"--", "--", "--", "--", "--", "--", "--", "--", "--", "xxx"}
	var expect = [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 30}
	result := calc(score1)
	fmt.Println("")
	assert.Equal(t, expect, result, "Total should be 30")
}

func TestCaseLastSpare(t *testing.T) {
	var score1 = [10]string{"--", "--", "--", "--", "--", "--", "--", "--", "--", "1/9"}
	var expect = [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 19}
	result := calc(score1)
	fmt.Println("")
	assert.Equal(t, expect, result, "Total should be 19")
}

func TestCaseLastStrike(t *testing.T) {
	var score1 = [10]string{"--", "--", "--", "--", "--", "--", "--", "--", "--", "xx5"}
	var expect = [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 25}
	result := calc(score1)
	fmt.Println("")
	assert.Equal(t, expect, result, "Total should be 25")
}

func TestCaseLastNumber(t *testing.T) {
	var score1 = [10]string{"--", "--", "--", "--", "--", "--", "--", "--", "--", "53"}
	var expect = [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 8}
	result := calc(score1)
	fmt.Println("")
	assert.Equal(t, expect, result, "Total should be 8")
}

func TestCaseFrame9Strike(t *testing.T) {
	var score1 = [10]string{"--", "--", "--", "--", "--", "--", "--", "--", "x-", "71"}
	var expect = [10]int{0, 0, 0, 0, 0, 0, 0, 0, 18, 8}
	result := calc(score1)
	fmt.Println("")
	assert.Equal(t, expect, result, "Total should be 26")
}

func TestCaseFrame9Spare(t *testing.T) {
	var score1 = [10]string{"--", "--", "--", "--", "--", "--", "--", "--", "1/", "71"}
	var expect = [10]int{0, 0, 0, 0, 0, 0, 0, 0, 17, 8}
	result := calc(score1)
	fmt.Println("")
	assert.Equal(t, expect, result, "Total should be 25")
}

func TestCaseSpareStrike(t *testing.T) {
	var score1 = [10]string{"1/", "x-", "5-", "--", "--", "--", "--", "--", "--", "--"}
	var expect = [10]int{20, 15, 5, 0, 0, 0, 0, 0, 0, 0}
	result := calc(score1)
	fmt.Println("")
	assert.Equal(t, expect, result, "Total should be 40")
}

func TestCaseStrikeSpare(t *testing.T) {
	var score1 = [10]string{"x-", "1/", "52", "--", "--", "--", "--", "--", "--", "--"}
	var expect = [10]int{20, 15, 7, 0, 0, 0, 0, 0, 0, 0}
	result := calc(score1)
	fmt.Println("")
	assert.Equal(t, expect, result, "Total should be 42")
}
