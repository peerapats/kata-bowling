package main

import "testing"
import "github.com/stretchr/testify/assert"
import "fmt"

func TestCaseZero(t *testing.T) {
	var score1 = [10]string{"--", "--", "--", "--", "--", "--", "--", "--", "--", "--"}
	var expect = [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	result := calc(score1)
	fmt.Println(result)
	fmt.Println("")
	assert.Equal(t, expect, result, "Error")
}

func TestCaseSuperStrike(t *testing.T) {
	var score1 = [10]string{"x-", "x-", "x-", "x-", "x-", "x-", "x-", "x-", "x-", "xxx"}
	var expect = [10]int{30, 30, 30, 30, 30, 30, 30, 30, 30, 30}
	result := calc(score1)
	fmt.Println(result)
	fmt.Println("")
	assert.Equal(t, expect, result, "Error")
}

func TestCase1(t *testing.T) {
	var score1 = [10]string{"62", "--", "--", "--", "--", "--", "--", "--", "--", "--"}
	var expect = [10]int{8, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	result := calc(score1)
	fmt.Println(result)
	fmt.Println("")
	assert.Equal(t, expect, result, "Error")
}

func TestCase2(t *testing.T) {
	var score1 = [10]string{"4/", "52", "--", "--", "--", "--", "--", "--", "--", "--"}
	var expect = [10]int{15, 7, 0, 0, 0, 0, 0, 0, 0, 0}
	result := calc(score1)
	fmt.Println(result)
	fmt.Println("")
	assert.Equal(t, expect, result, "Error")
}

func TestCase3(t *testing.T) {
	var score1 = [10]string{"62", "4/", "x-", "4/", "x-", "x-", "x-", "x-", "x-", "4/5"}
	var expect = [10]int{8, 20, 20, 20, 30, 30, 30, 24, 20, 15}
	result := calc(score1)
	fmt.Println(result)
	fmt.Println("")
	assert.Equal(t, expect, result, "Error")
}

func TestCase4(t *testing.T) {
	var score1 = [10]string{"--", "4/", "x-", "4/", "x-", "x-", "x-", "x-", "x-", "4/5"}
	var expect = [10]int{0, 20, 20, 20, 30, 30, 30, 24, 20, 15}
	result := calc(score1)
	fmt.Println(result)
	fmt.Println("")
	assert.Equal(t, expect, result, "Error")
}
