package main

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
}
