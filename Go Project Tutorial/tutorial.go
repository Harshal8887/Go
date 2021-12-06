package main

import "fmt"

func main() {
	fmt.Println("!!Welcome to the Quiz game!!")
	var name string
	fmt.Printf("Enter yor name: \n")
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the game!!\n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yay you can Play!")
	} else {
		fmt.Println("You cannot play!")
		return
	}

	score := 0
	num_ques := 2

	fmt.Printf("Which is beter, the RTX 3080 or RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)
	if answer+answer2 == "RTX 3090" || answer+answer2 == "rtx 3090" {
		fmt.Println("Correct!")
		score += 1
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("How many cores does the Ryzen 9 3090x have?")
	var cores uint
	fmt.Scan(&cores)
	if cores == 12 {
		fmt.Println("Correct!")
		score++

	} else {
		fmt.Println("Incorrect!")
	}
	fmt.Printf("You Scored %v out of %v.\n", score, num_ques)
	percent := (float64(score) / float64(num_ques)) * 100
	fmt.Printf("You scored : %v%%, percent\n", percent)

	//fmt.Println(answer, answer2)
	//fmt.Println(age >= 10)
	//var name string = "Tim"
	// int uint{+ve} float64 bool
	// name := "Tim"
	// age := 21
	// fmt.Printf("Hello %v, you re %v? ", name, age) %f %d
	// else if answer+answer2 == "rtx 3090" {
	// 	fmt.Println("Correct!")
	// 	score++	}

}
