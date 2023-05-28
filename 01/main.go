package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var maxSums = []int{}

func main() {
	f, err := os.Open("input.in")
	if err != nil {
		panic("omg")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	runningSum := 0

	for scanner.Scan() {
		text := scanner.Text()
		value, _ := strconv.Atoi(text)
		runningSum += value

		if len(text) == 0 {
			maxSums = append(maxSums, runningSum)

			runningSum = 0
		}
	}
	// sort maxSums
	sort.Slice(maxSums, func(i, j int) bool {
		return maxSums[i] > maxSums[j]
	})

	fmt.Println(maxSums[0] + maxSums[1] + maxSums[2])
}
