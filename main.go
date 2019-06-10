package main

import "fmt"

func main() {
	scores := [5]string{"9", "0", "9", "1", "10"}

	for index, name := range scores {



		fmt.Println(index, name)
	}
}
