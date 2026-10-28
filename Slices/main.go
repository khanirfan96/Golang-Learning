package main

import (
	"fmt"
)

func main() {

	// arr := [5]int{10, 20, 30, 40, 50}

	// s := arr[1:4]

	// fmt.Println("Simple ----", arr)
	// fmt.Println("Slice ----", s)
	// fmt.Println("Length ----", len(s))
	// fmt.Println("Capacity ----", cap(s))

	// Through make keyword

	// s := make([]int, 3, 5)

	// fmt.Println("Slice ----", s)
	// fmt.Println("Length ----", len(s))
	// fmt.Println("Capacity ----", cap(s))

	// s := []int{1, 2, 3}
	// s = append(s, 5)
	// fmt.Println("append", s)

	// s := make([]int, 3, 3)

	// fmt.Println("Before append ----", s, len(s), cap(s))

	// s = append(s, 10)
	// fmt.Println("Before append ----", s, len(s), cap(s))

	// Delete slice

	s := []int{10, 20, 30, 40, 50}

	fmt.Println("Old Slice ----", s)

	i := 2
	s = append(s[:i], s[i+1:]...)
	fmt.Println("New Slice ----", s)

}
