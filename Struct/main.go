package main

import (
	"encoding/json"
	"fmt"
)

// Creation of struct

type User struct {
	ID   int
	Name string
	Age  int
	Address
}

type Address struct {
	City  string
	State string
}

// Json converter Struct
type Convert struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `josn:"age"`
	Password string `json:"-"`
}

// Value Receiver
func (u User) Greet() {
	fmt.Println("Hiii", u.Name)
}

// Pointer Receiver
func (u *User) Update(name string) {
	u.Name = name
}

func main() {

	// user := User{
	// 	ID:   1,
	// 	Name: "Irfan Khan",
	// 	Age:  30,
	// }
	// fmt.Println("User Struct----", user)

	// Intialization with new keyword
	// user2 := new(User)
	// user2.ID = 1
	// user2.Name = "Irfan"
	// user2.Age = 30

	// fmt.Println("New Struct---", user2)

	// Intialization with pointer
	// user3 := &User{
	// 	ID:   1,
	// 	Name: "Irfan Khan",
	// 	Age:  30,
	// }

	// fmt.Println("Pointer Struct---", user3)

	//For Value & Pointer Receiver syntax
	// user := User{
	// 	ID:   1,
	// 	Name: "Irfan Khan",
	// 	Age:  30,
	// }

	// user.Greet()

	// user.Update("Babayaga")
	// fmt.Println("Receiver User", user.Name)

	// Nested/Composition Struct
	// user := User{
	// 	ID:   1,
	// 	Name: "Irfan Khan",
	// 	Age:  30,
	// 	Address: Address{
	// 		City:  "Allahabad",
	// 		State: "U.P",
	// 	},
	// }
	// fmt.Println("Nested Struct", user)

	// Converting Struct into JSON

	user := Convert{
		ID:       1,
		Name:     "Irfan Khan",
		Age:      30,
		Password: "Secret",
	}

	jsonDate, err := json.Marshal(user)

	if err != nil {
		fmt.Println("Error")
		return
	}

	fmt.Println("Json Data", string(jsonDate))
}
