package main

import(
"os"
"fmt"
)


func main() {
	for i,value := range os.Args[0:]{
		fmt.Sprintf("Index = %d and Value = %v", i, value)
		fmt.Println("Printing")
	}
}