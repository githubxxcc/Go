package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	strSlice := os.Args[1:]
	var slice []int
	for i := range strSlice {
		x, ok := strconv.Atoi(strSlice[i])

		if ok == nil {
			slice = append(slice, x)

		}
	}
	fmt.Println(slice)

	var a *tree

	for i := range slice {
		a = add(a, slice[i])
	}
	fmt.Println(slice)

	appendValues(slice[:0], a)

	fmt.Println(slice)
}

type tree struct {
	left, right *tree
	value       int
}

func add(a *tree, x int) *tree {
	if a == nil {
		return &tree{value: x}
	}

	if x < a.value {
		a.left = add(a.left, x)
	} else {
		a.right = add(a.right, x)
	}

	return a
}

func appendValues(values []int, a *tree) []int {
	if a != nil {
		values = appendValues(values, a.left)
		values = append(values, a.value)
		values = appendValues(values, a.right)
	}

	return values
}
