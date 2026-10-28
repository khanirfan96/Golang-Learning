package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// var name string
	// var age int

	// fmt.Print("Enter your name: ")
	// fmt.Scan(&name)

	// fmt.Print("Enter your age: ")
	// fmt.Scan(&age)

	// fmt.Printf("My name is %s and my age is %d\n", name, age)

	readValue := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your Full Name----")
	name, _ := readValue.ReadString('\n')
	fmt.Println("Your name is----", name)

}
