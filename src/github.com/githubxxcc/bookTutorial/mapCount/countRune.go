package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	var utflens [utf8.UTFMax + 1]int
	invalid := 0

	input := bufio.NewReader(os.Stdin)

	for {
		r, n, err := input.ReadRune()

		if err == io.EOF {
			break
		}

		// fmt.Println("1")

		if err != nil {
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		// fmt.Println("3")
		counts[r]++
		utflens[n]++
	}

	fmt.Printf("%v", counts)
}
