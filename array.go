package main

import "fmt"

func ArrayExample() {
	arrayExample := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(arrayExample[1:4]) // [1 5 3]
	fmt.Println(arrayExample[1:])  // [1 5 3 6 4]
	fmt.Println(arrayExample[:4])  // [7 1 5 3]
	fmt.Println(arrayExample[:])   // [7 1 5 3 6 4]
	fmt.Println(arrayExample[0])   // 7
	featuredExample := arrayExample[1:4]
	featuredExample[0] = 100
	highlightedExample := arrayExample[1:4:4]
	highlightedExample[0] = 1000
	appendedExample := append(highlightedExample, 8)
	fmt.Println(arrayExample)                               // [7 100 5 3 6 4]
	fmt.Println(featuredExample)                            // [100 5 3]
	fmt.Println(appendedExample)                            // [1000 5 3]
	fmt.Println(len(appendedExample), cap(appendedExample)) // 3
}
