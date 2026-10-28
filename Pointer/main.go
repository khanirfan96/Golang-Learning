package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func updateAge(u *User) {
	u.Age = 100
}

func increment(n *int) {
	*n = *n + 1
}

func main() {

	// x := 10
	// ptr := &x
	// fmt.Println("Value of x----", x)
	// fmt.Println("Address of x----", &x)
	// fmt.Println("Value stored in Pointer----", ptr)
	// fmt.Println("Value at ptr----", *ptr)

	// Pointer to struct

	// user := User{
	// 	Name: "Irfan Khan",
	// 	Age:  30,
	// }
	// updateAge(&user)
	// fmt.Println("Pointer with struct", user)

	// Pointer as Function Parameter

	x := 10
	increment(&x)

	fmt.Println("Pointer Function", x)

}
