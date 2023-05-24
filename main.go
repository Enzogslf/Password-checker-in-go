package main

import "fmt"

func main() {
	var password string = "891"
	var input string
	var isCorrect bool
	// var notCorrect bool

	for {

		fmt.Scan(&input)
		if input == password {
			isCorrect = true

		} else {
			fmt.Println("thats not right!")
		}

		if isCorrect == true {
			fmt.Println("thats correct!")
			break
		}
	}
}
