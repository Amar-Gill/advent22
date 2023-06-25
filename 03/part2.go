package main

import (
	"bufio"
	"fmt"
	"os"
)

const GROUP_SIZE = 3

func main() {
	priorities := []int{}
	f, err := os.Open("input.in")
	if err != nil {
		panic("omg")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	rucksacks := []string{}
	for scanner.Scan() {
		text := scanner.Text()
		rucksacks = append(rucksacks, text)
		if len(rucksacks) == GROUP_SIZE {
			runeSet := initializeRuneSet(rucksacks[0])
			for i := 1; i < GROUP_SIZE; i++ {
				runeSet = updateRuneSet(rucksacks[i], runeSet, i)
			}
			rucksacks = rucksacks[:0]
			// should be map length 1 at this point
			for k := range runeSet {
				priorities = append(priorities, convertRuneToPriority(k))
			}
		}
	}

	sum := 0
	for _, v := range priorities {
		sum += v
	}
	fmt.Println(sum) // 2716
}

func initializeRuneSet(s string) map[rune]int {
	ret := map[rune]int{}

	for _, v := range s {
		if _, ok := ret[v]; !ok {
			ret[v] = 0
		}
	}
	return ret
}

func updateRuneSet(s string, m map[rune]int, iteration int) map[rune]int {
	ret := m

	for _, v := range s {
		if _, ok := ret[v]; !ok {
			delete(ret, v)
		} else if ret[v] == iteration {
			continue
		} else {
			ret[v]++
		}
	}

	for k, v := range ret {
		if v != iteration {
			delete(ret, k)
		}
	}
	return ret
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
