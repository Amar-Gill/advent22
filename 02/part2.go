package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	opponent       = map[string]string{"A": "rock", "B": "paper", "C": "scissors"}
	gameRules      = map[string]map[string]string{"rock": {"win": "paper", "lose": "scissors", "draw": "rock"}, "paper": {"win": "scissors", "lose": "rock", "draw": "paper"}, "scissors": {"win": "rock", "lose": "paper", "draw": "scissors"}}
	myScore        = 0
	resultToObtain = map[string]string{"X": "lose", "Y": "draw", "Z": "win"}
)

func main() {
	f, err := os.Open("input.in")
	if err != nil {
		panic("omg")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			break
		}

		choices := strings.Split(text, "")
		opChoice := opponent[choices[0]]
		requiredResult := resultToObtain[choices[2]]
		myScore += getGameResult(opChoice, requiredResult)
	}

	fmt.Println(myScore)
}

func getGameResult(opChoice string, requiredResult string) int {
	result := 0
	myChoice := gameRules[opChoice][requiredResult]

	fmt.Printf("Op choice is: %s, required result is: %s, so I play: %s\n", opChoice, requiredResult, myChoice)

	switch requiredResult {
	case "win":
		result += 6
	case "draw":
		result += 3
	}

	switch myChoice {
	case "rock":
		result += 1
	case "paper":
		result += 2
	case "scissors":
		result += 3
	}

	return result
}
