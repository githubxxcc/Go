package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	seen := make(map[string]bool)

	input := bufio.Scanner(os.Stdin)

	for input.Scan(){
		s := input.Text()
		if !seen[s]{
			seen[s] = true
		}
	}

	if err:= input.Err(); err != nil{
		fmt.Fprintf(os.Stderr, %s, err)
		os.Exit(1)
	}

}
