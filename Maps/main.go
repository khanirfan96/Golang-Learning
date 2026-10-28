package main

import "fmt"

func main() {

	// Cration of maps

	// m := make(map[string]int)
	// m["age"] = 30
	// fmt.Println("maps---", m)

	// m := map[string]string{
	// 	"name": "Irfan",
	// 	"age":  "30",
	// }
	// fmt.Println("maps----", m)

	m := make(map[string]int)
	m["age"] = 30
	m["salary"] = 10
	m["num"] = 91
	m["bp"] = 120
	// fmt.Println("maps---", m["age"])  // 30
	// fmt.Println("maps---", m["jiii"]) //0

	// if val, ok := m["val"]; ok {
	// 	fmt.Println("Key Exist", val)
	// } else {
	// 	fmt.Println("Not Exist")
	// }

	// for key, value := range m {
	// 	fmt.Println(key, "   ", value)
	// }

	// for key := range m {
	// 	fmt.Println(key)
	// }

	// for _, value := range m {
	// 	fmt.Println(value)
	// }

	// delete(m, "num")
	// fmt.Println("delete a key", m)

	type User struct {
		Name string
		Age  int
	}

	users := make(map[int]User)
	users[1] = User{"Irfan", 30}
	fmt.Println("users------", users)

}
