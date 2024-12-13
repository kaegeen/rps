package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func getComputerChoice() string {
	choices := []string{"rock", "paper", "scissors"}
	return choices[rand.Intn(len(choices))]
}

func determineWinner(player, computer string) string {
	if player == computer {
		return "It's a tie!"
	}

	switch player {
	case "rock":
		if computer == "scissors" {
			return "You win! Rock smashes scissors."
		}
	case "paper":
		if computer == "rock" {
			return "You win! Paper covers rock."
		}
	case "scissors":
		if computer == "paper" {
			return "You win! Scissors cut paper."
		}
	}

	return fmt.Sprintf("You lose! %s beats %s.", strings.Title(computer), strings.Title(player))
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed random number generator

	fmt.Println("Rock-Paper-Scissors Game")
	fmt.Println("========================")
	fmt.Println("Enter 'rock', 'paper', or 'scissors' to play. Type 'exit' to quit.")

	for {
		// Get player's choice
		fmt.Print("Your choice: ")
		var playerChoice string
		fmt.Scan(&playerChoice)
		playerChoice = strings.ToLower(strings.TrimSpace(playerChoice))

		// Exit condition
		if playerChoice == "exit" {
			fmt.Println("Thanks for playing! Goodbye!")
			break
		}

		// Validate player's choice
		if playerChoice != "rock" && playerChoice != "paper" && playerChoice != "scissors" {
			fmt.Println("Invalid choice. Please enter 'rock', 'paper', or 'scissors'.")
			continue
		}

		// Get computer's choice
		computerChoice := getComputerChoice()
		fmt.Printf("Computer chose: %s\n", computerChoice)

		// Determine and display the winner
		result := determineWinner(playerChoice, computerChoice)
		fmt.Println(result)
	}
}
