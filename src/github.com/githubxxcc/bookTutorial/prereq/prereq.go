package main

import (
	"fmt"
	"sort"
)

var prereq = map[string][]string{
	"complex algorithm": {
		"data structure",
		"intro to algorithm",
	},
	"graph": {
		"intro to algorithm",
		"data structure",
	},
	"data structure": {"intro to algorithm"},
}

func main() {
	for i, value := range topoSort(prereq) {
		fmt.Printf("%d : %s \n", i+1, value)
	}
}

func topoSort(m map[string][]string) []string {
	var list []string
	seen := make(map[string]bool)

	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {

			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				//do something
				list = append(list, item)
			}
		}
	}

	var keys []string

	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	visitAll(keys)

	return list
}
