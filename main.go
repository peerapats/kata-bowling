package main

import (
	"fmt"
	"strconv"
	"strings"
)

func cal(key int, val string, input [10]string) int64 {
	frameInput := strings.Split(val, "")
	frameInputInt := int64(0)
	for _, v := range frameInput {
		if v == "x" {
			frameInputInt = int64(10)
			j := string(input[key+1][0])
			k := string(input[key+2][0])
			if j == "x" {
				frameInputInt += int64(10)
				if k == "x" {
					frameInputInt += int64(10)
				} else {
					nextFrame1, _ := strconv.ParseInt(string(input[key+1][0]), 10, 64)
					frameInputInt += nextFrame1
				}
			} else {
				nextFrame1, _ := strconv.ParseInt(string(input[key+1][0]), 10, 64)
				frameInputInt += nextFrame1
			}
			break
		}
		if v == "/" {
			frameInputInt = int64(10)
			j := string(input[key+1][0])
			nextFrame1 := int64(0)
			if j == "x" {
				nextFrame1 = int64(10)
			} else {
				nextFrame1, _ = strconv.ParseInt(string(input[key+1][0]), 10, 64)
			}
			frameInputInt += nextFrame1
			break
		}
		i, _ := strconv.ParseInt(v, 10, 64)
		frameInputInt += i
	}
	return frameInputInt
}

func main() {
	var games = [10]string{"2/", "5/", "x", "x", "x", "10", "53", "63", "27", "11"}
	fmt.Println("Input => ", games)
	totalScore := int64(0)
	for key, value := range games {
		frameInputInt := cal(key, value, games)
		totalScore += frameInputInt
		fmt.Println("Frame", key+1, " = ", frameInputInt)
	}
	fmt.Println("Total: ", totalScore)
}
