package main

import (
	"fmt"
)

func main() {
	list, err := NewList(3, 50)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(list.Length()) // 3
	list.Add(1)
	list.Add(2)
	list.Add(8)
	// fmt.Println(list.Get(0)) // 50
	// fmt.Println(list.Get(1)) // 51
	// fmt.Println(list.Get(4)) // 2
	// fmt.Println(list.Get(5)) // 8
	// list.PrintAll()
	list.Remove(0)
	// list.PrintAll()
	list.Remove(2)
	list.PrintAll()            // [51, 52, 2, 8]
	fmt.Println(list.Length()) // 4
}
