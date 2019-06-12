package main

import (
	"fmt"
	"strings"
	"strconv"
)

var score1 = [10]string{"x-", "x-", "x-", "x-", "x-", "x-", "x-", "x-", "x-", "12"}
var score2 = [10]string{"4/", "5/", "x-", "-/", "--", "--", "--", "--", "--", "78"}
var score3 = [10]string{"-/", "6-", "x-", "71", "8-", "x-", "6/", "7/", "-/", "33"}

var score = [10]string{"x-", "1/", "12", "12", "12", "12", "12", "12", "12", "12"}

func main() {
	// fmt.Println("Hello Golang!")
	calc(score3)
}

func calc(scores [10]string) {
	fmt.Println(scores)
	// sum := 0
	for i := range scores {
		// fmt.Println(scores[i])
		s := strings.Split(scores[i], "")
		// fmt.Println(s)
		throwing(s, scores, i)
	}
}

// func calLast(s []string, scores [10]string, index int) {

// }

func throwing(s []string, scores [10]string, index int) {
	// fmt.Println(s)
	sum := 0
	switch s[0] {
		case "-":
			sum = 0
		case "x":
			s2 := strings.Split(scores[index + 1], "")
			sum = calculateStrike(s2, scores, index)
		case "1", "2", "3", "4", "5", "6", "7", "8", "9" :
			sum = convertStringToInt(s[0])
	}

	switch s[1] {
		case "-":
			sum = sum + 0
		case "1", "2", "3", "4", "5", "6", "7", "8", "9" :
			sum = sum + convertStringToInt(s[1])
		case "/":
			s2 := strings.Split(scores[index + 1], "")
			sum = calSpare(s2)
	}	

	fmt.Println(sum)
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
		s3 := strings.Split(scores[index + 2], "")
		if s3[0] == "x" {
			return 30
		} else {
			return 20 + convertStringToInt(s3[0])
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
