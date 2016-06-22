package main

import (
	"bufio"
	"io"
	"os"
	"unicode"
)

func main() {
	digitCounts := make(map[rune]int)
	letterCounts := make(map[rune]int)

	invalid := 0

	input := bufio.NewReader(os.Stdin)

	for {
		x, n, err := input.ReadRune()

		if err == io.EOF {
			break
		}

		if err != nil {
			os.Exit(1)
		}

		if x == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		if unicode.IsLetter(x) {
			letterCounts[x]++
		}

		if unicode.IsDigit(x) {
			digitCounts[x]++
		}
	}
}
