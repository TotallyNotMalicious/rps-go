package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		choice1 string
		choice2 string
	)

	fmt.Println("Player 1\nChoose an option\n\nRock\nPaper\nScissors")
	fmt.Print("\nYour choice: ")
	fmt.Scan(&choice1)
	option1 := strings.ToLower(choice1)

	fmt.Println("Player 2\nChoose an option\n\nRock\nPaper\nScissors")
	fmt.Print("\nYour choice: ")
	fmt.Scan(&choice2)
	option2 := strings.ToLower(choice2)

	switch { // create a switch with all of our possible combinations
	case option1 == option2:
		fmt.Println("\nYour choices match, you must choose 2 different choices.")
	case option1 == "rock" && option2 == "paper":
		fmt.Println("\nPlayer 2 wins!")
	case option1 == "paper" && option2 == "rock":
		fmt.Println("\nPlayer 1 wins!")
	case option1 == "rock" && option2 == "scissors":
		fmt.Println("\nPlayer 1 wins!")
	case option1 == "scissors" && option2 == "rock":
		fmt.Println("\nPlayer 2 wins!")
	case option1 == "paper" && option2 == "rock":
		fmt.Println("\nPlayer 1 wins!")
	case option1 == "rock" && option2 == "paper":
		fmt.Println("\nPlayer 2 wins!")
	case option1 == "scissors" && option2 == "paper":
		fmt.Println("\nPlayer 1 wins!")
	case option1 == "paper" && option2 == "scissors":
		fmt.Println("\nPlayer 2 wins!")
	default:
		fmt.Println("\nInvalid choice, please choose either rock, paper, or scissors.")
	}
}
