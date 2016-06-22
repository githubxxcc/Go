package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	fmt.Println(addComma(os.Args[1]))
}

func addComma(s string) string {
	var buf bytes.Buffer

	for i := 0; i < len(s); i++ {
		if (len(s)-i)%3 == 0 && i > 0 {
			buf.WriteByte(',')
		}

		buf.WriteByte(s[i])
	}

	return buf.String()
}
