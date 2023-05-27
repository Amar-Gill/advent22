package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.in")
	if err != nil {
		panic("omg")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	max := 0
	runningSum := 0

	for scanner.Scan() {
		text := scanner.Text()
		value, _ := strconv.Atoi(text)
		runningSum += value

		if len(text) == 0 {
			if runningSum > max {
				max = runningSum
			}

			runningSum = 0
		}
	}

	fmt.Println(max)
}
