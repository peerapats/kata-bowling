package main

import (
	"fmt"
	"strconv"
	"strings"
)

var score0 = [10]string{"x-", "x-", "x-", "x-", "x-", "x-", "x-", "x-", "x-", "xxx"}
var score1 = [10]string{"x-", "x-", "x-", "x-", "x-", "x-", "x-", "x-", "x-", "12"}
var score2 = [10]string{"4/", "5/", "x-", "-/", "--", "--", "--", "--", "--", "78"}
var score3 = [10]string{"-/", "6-", "x-", "71", "8-", "x-", "6/", "7/", "-/", "33"}

var score = [10]string{"x-", "1/", "12", "12", "12", "12", "12", "12", "12", "12"}

func main() {
	// fmt.Println("Hello Golang!")
	result := calc(score0)
	fmt.Println(result)
}

func calc(scores [10]string) [10]int {
	fmt.Println(scores)
	var result = [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := range scores {
		// fmt.Println(scores[i])
		s := strings.Split(scores[i], "")
		// fmt.Println(s)

		if i == 9 {
			result[i] = throwLast(s, scores, i)
		} else {
			result[i] = throwing(s, scores, i)
		}
	}

	return result
}

func throwLast(s []string, scores [10]string, index int) int {
	return 0
}

func throwing(s []string, scores [10]string, index int) int {
	// fmt.Println(s)
	sum := 0
	switch s[0] {
	case "-":
		sum = 0
	case "x":
		s2 := strings.Split(scores[index+1], "")
		sum = calculateStrike(s2, scores, index)
	case "1", "2", "3", "4", "5", "6", "7", "8", "9":
		sum = convertStringToInt(s[0])
	}

	switch s[1] {
	case "-":
		sum = sum + 0
	case "1", "2", "3", "4", "5", "6", "7", "8", "9":
		sum = sum + convertStringToInt(s[1])
	case "/":
		s2 := strings.Split(scores[index+1], "")
		sum = calSpare(s2)
	}

	return sum
}

func calSpare(s2 []string) int {
	if s2[0] == "x" {
		return 20
	} else {
		return 10 + convertStringToInt(s2[0])
	}
}

func calculateStrike(s2 []string, scores [10]string, index int) int {
	if s2[0] == "x" {
		if index != 8 {
			s3 := strings.Split(scores[index+2], "")
			if s3[0] == "x" {
				return 30
			} else {
				return 20 + convertStringToInt(s3[0])
			}
		} else {
			return 20 + convertStringToInt(s2[1])
		}
	} else {
		if s2[1] == "/" {
			return 20
		} else {
			return 10 + convertStringToInt(s2[0]) + convertStringToInt(s2[1])
		}
	}
}

func convertStringToInt(str string) int {
	if str == "-" {
		return 0
	}

	if str == "x" {
		return 10
	}

	i1, err := strconv.Atoi(str)
	if err == nil {
		return i1
	} else {
		return 0
	}
}
