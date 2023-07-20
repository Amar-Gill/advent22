package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("input.in")
	if err != nil {
		panic("omg")
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		fmt.Println(r)
	}
}
