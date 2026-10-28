package main

import (
	"fmt"
)

func main() {

	// var arr [5]int

	// arr[0] = 100
	// arr[1] = 200

	// fmt.Println("Arrays---", arr)

	arr := [4]int{2, 3, 4, 5}
	arr1 := [4]int{2, 3, 4}
	fmt.Println("arrr----", arr)
	fmt.Println("Length----", len(arr))

	for i := 0; i < len(arr); i++ {
		fmt.Println("Normal for loop ----", arr[i])
	}

	for index, value := range arr {
		fmt.Println(index, "Range loop ------", value)
	}

	fmt.Println("Compare array ----", arr == arr1)

	// To copy an array
	arr2 := arr
	fmt.Println("Arrraayyyy2222------", arr2)
}
