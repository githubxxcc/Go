package main

import ("fmt"
"strings"
"github.com/githubxxcc/stringutil"
)


// Comments 

func main(){
	
	favNum3 := [5]float64 {1,2,3,4,5}

	for i, value := range favNum3{ 
		fmt.Println(value,i)
	}

	numSlice := []int {5,4,3,2,1}

	numSlice3 := make([]int, 7,10)

	copy(numSlice3, numSlice)

	fmt.Println(numSlice3[:])

	numSlice3 = append(numSlice3, -1)

	fmt.Println(numSlice3[:])

	presAge := make(map[string] int)

	presAge["ME"] = 23
	presAge["HER"] = 23
	delete(presAge, "HdsER")

	fmt.Println(presAge)

	someString := "HelloWorld"

	fmt.Println(strings.Contains(someString, "hi"))

	fmt.Println(strings.Index(someString, "lo"))

	fmt.Println(strings.Count(someString, "l"))

	fmt.Println(strings.Replace(someString, "l", "x", 2))

	fmt.Println(stringutil.Reverse("!oG, olleH"))
	
}