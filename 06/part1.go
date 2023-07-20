package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var markerLength = 4

func main() {
	f, err := os.Open("input.in")
	if err != nil {
		panic("omg")
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	sequence := []rune{}
	charCount := -1
	for {
		charCount++

		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}

		if len(sequence) < markerLength {
			sequence = append(sequence, r)
			continue
		}

		runeSet := buildRuneSet(sequence)

		if len(runeSet) == markerLength {
			// means we have "markerLength" unique runes in the set, so we found the marker
			break
		}

		sequence = sequence[1:]
		sequence = append(sequence, r)
	}

	fmt.Println(charCount)
}

func buildRuneSet(runes []rune) map[rune]int {
	ret := map[rune]int{}

	for _, r := range runes {
		if _, ok := ret[r]; !ok {
			ret[r] = 1
			continue
		}
		ret[r]++
	}

	return ret
}
