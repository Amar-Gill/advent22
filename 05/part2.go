package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	stack1 := []rune{'N', 'R', 'G', 'P'}
	stack2 := []rune{'J', 'T', 'B', 'L', 'F', 'G', 'D', 'C'}
	stack3 := []rune{'M', 'S', 'V'}
	stack4 := []rune{'L', 'S', 'R', 'C', 'Z', 'P'}
	stack5 := []rune{'P', 'S', 'L', 'V', 'C', 'W', 'D', 'Q'}
	stack6 := []rune{'C', 'T', 'N', 'W', 'D', 'M', 'S'}
	stack7 := []rune{'H', 'D', 'G', 'W', 'P'}
	stack8 := []rune{'Z', 'L', 'P', 'H', 'S', 'C', 'M', 'V'}
	stack9 := []rune{'R', 'P', 'F', 'L', 'W', 'G', 'Z'}

	stacks := make(map[int][]rune, 9)
	stacks[1] = stack1
	stacks[2] = stack2
	stacks[3] = stack3
	stacks[4] = stack4
	stacks[5] = stack5
	stacks[6] = stack6
	stacks[7] = stack7
	stacks[8] = stack8
	stacks[9] = stack9

	f, err := os.Open("instructions.in")
	if err != nil {
		panic("omg")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text := scanner.Text()
		tokens := strings.Split(text, " ")
		numCratesToMove, _ := strconv.Atoi(tokens[1])
		fromStack, _ := strconv.Atoi(tokens[3])
		toStack, _ := strconv.Atoi(tokens[5])

		newFrom, newTo := moveCrates(stacks[fromStack], stacks[toStack], numCratesToMove)

		stacks[fromStack] = newFrom
		stacks[toStack] = newTo
	}

	topCrates := []rune{}
	for i := 1; i < 10; i++ {
		stack := stacks[i]
		topCrates = append(topCrates, stack[len(stack)-1])
	}
	fmt.Println(string(topCrates))
}

func moveCrates(fromStack []rune, toStack []rune, numCratesToMove int) ([]rune, []rune) {
	newFrom := fromStack
	newTo := toStack

	for i := 0; i < numCratesToMove; i++ {

		pop, s := pop(newFrom)
		newFrom = s
		newTo = append(newTo, pop)
	}
	return newFrom, newTo
}

// https://blog.toshima.ru/2021/10/18/go-slice-pop-shift.html
func pop[T any](s []T) (T, []T) {
	pop, s1 := s[len(s)-1], s[:len(s)-1]
	return pop, s1
}

// returns two slices
// first is	split from beggining to index -1
// second is from index to end
func splitSlice[T any](s []T, index int) ([]T, []T) {
	s1, s2 := s[:index], s[index:]
	return s1, s2
}
