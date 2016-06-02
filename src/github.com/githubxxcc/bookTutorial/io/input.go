package main

import("fmt" 
"os" 
"bufio"
)

func main() {
	counts := make(map[string]int)
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan(){
		counts[sc.Text()]++
	}

	for i, value := range counts{
		fmt.Printf("String : %s, Counts : %d \n", i, value)
	}


}