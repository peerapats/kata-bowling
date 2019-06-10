package main

import "fmt"

func main() {

	input := [10]string{
		"x,-",
		"x,-",
		"x,-",
		"x,-",
		"x,-",
		"x,-",
		"x,-",
		"x,-",
		"x,-",
		"xxx",
	}

	// for index, name := range input {

	// }

	fmt.Println(input)

	scores := [5]string{"9", "0", "9", "1", "10"}

	for index, name := range scores {
		fmt.Println(index, name)
	}

}
