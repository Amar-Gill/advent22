package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	priorities := []int{}
	f, err := os.Open("input.in")
	if err != nil {
		panic("omg")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text := scanner.Text()
		length := len(text)
		compartment1 := text[0 : length/2]
		compartment2 := text[length/2 : length]

		for _, v := range compartment1 {
			containsRune := strings.ContainsRune(compartment2, v)

			if containsRune {
				priorities = append(priorities, convertRuneToPriority(v))
				break
			}
		}
	}

	sum := 0
	for _, v := range priorities {
		sum += v
	}
	fmt.Println(sum)
}

// Lowercase item types a through z have priorities 1 through 26.
// Uppercase item types A through Z have priorities 27 through 52.
// Rune Values: A = 65, Z = 90, a = 97, z = 122
func convertRuneToPriority(r rune) int {
	const minLowerCase = 'a'
	const maxLowerCase = 'z'

	isLowerCaseLetter := (r >= minLowerCase && r <= maxLowerCase)

	if isLowerCaseLetter {
		return int(r) - 96
	}
	return int(r) - 38
}
