package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	opponent  = map[string]string{"A": "rock", "B": "paper", "C": "scissors"}
	me        = map[string]string{"X": "rock", "Y": "paper", "Z": "scissors"}
	gameRules = map[string]map[string]string{"rock": {"defeats": "scissors", "defeatedBy": "paper"}, "paper": {"defeats": "rock", "defeatedBy": "scissors"}, "scissors": {"defeats": "paper", "defeatedBy": "rock"}}
	myScore   = 0
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
		myChoice := me[choices[2]]
		myScore += getGameResult(opChoice, myChoice)
	}

	fmt.Println(myScore)
}

func getGameResult(opChoice string, myChoice string) int {
	result := 0

	if opChoice == myChoice {
		result += 3
	} else if gameRules[myChoice]["defeats"] == opChoice {
		result += 6
	} else {
		result += 0
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
