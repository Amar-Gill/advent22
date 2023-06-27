package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.in")
	if err != nil {
		panic("omg")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	numPairs := 0
	for scanner.Scan() {
		text := scanner.Text()
		vals := strings.Split(text, ",")
		elf1Range := strings.Split(vals[0], "-")
		elf2Range := strings.Split(vals[1], "-")

		elf1Min, _ := strconv.Atoi(elf1Range[0])
		elf1Max, _ := strconv.Atoi(elf1Range[1])

		elf2Min, _ := strconv.Atoi(elf2Range[0])
		elf2Max, _ := strconv.Atoi(elf2Range[1])

		if elf2Min > elf1Max {
			continue
		} else if elf1Min > elf2Max {
			continue
		} else {
			numPairs++
		}
	}
	fmt.Println(numPairs) // 808
}
