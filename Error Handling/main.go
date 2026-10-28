package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by 0")
	}

	ans := a / b
	return ans, nil
}

func findUser(id int) error {
	return fmt.Errorf("User with id %d not found", id)
}

func main() {

	result, err := divide(10, 2)

	if err != nil {
		fmt.Println("There is error---", err)
		return
	}

	fmt.Println("Result is--", result)

	c := findUser(2)
	fmt.Println("Result of C--", c)

}
