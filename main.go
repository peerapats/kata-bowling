package main

import (
	"fmt"
	"strings"
	"strconv"
)

var score1 = [10]string{"x", "x", "x", "x", "x", "x", "x", "x", "x", "xxx"}
var score2 = [10]string{"4/", "5/", "x", "-/", "--", "--", "--", "--", "--", "7/8"}
var score3 = [10]string{"-/", "6-", "x", "71", "8-", "x", "6/", "7/", "-/", "x4/"}

var score = [10]string{"x", "12", "12", "12", "12", "12", "12", "12", "12", "12"}

func main() {
	fmt.Println("Hello Golang!")
	calc(score)
}

func calc(scores [10]string) {
	// fmt.Println(scores)
	// sum := 0
	for i := range scores {
		// fmt.Println(scores[i])
		s := strings.Split(scores[i], "")
		// fmt.Println(s)
		throwing(s, scores, i)
	}
}

func throwing(s []string, scores [10]string, index int) {
	// fmt.Println(s)
	sum := 0
	switch s[0] {
		case "-":
			sum = 0
		case "x":
			s2 := strings.Split(scores[index + 1], "")
			sum = 10 + convertStringToInt(s2[0]) + convertStringToInt(s2[1])
		case "1", "2", "3", "4", "5", "6", "7", "8", "9" :
			sum = convertStringToInt(s[0])
	}

	fmt.Println(sum)
}

func convertStringToInt(str string) int {
	i1, err := strconv.Atoi(str)
	if err == nil {
		return i1
	} else {
		return 0
	}
}

// func summary()