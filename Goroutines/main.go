package main

import (
	"fmt"
	"sync"
)

// func Hello() {
// 	fmt.Println("Hello")
// }

// func task(name string) {
// 	for i := 1; i <= 4; i++ {
// 		fmt.Println(name, " ", i)
// 		time.Sleep(time.Millisecond * 2)
// 	}
// }

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Worker ---", id)
}

func main() {
	// 1st Example
	// go Hello()
	// time.Sleep(time.Second * 1) // Added because if i don't add this directly main func will exit,
	// // because go routine is not getting time to execute/start

	// 2nd Example
	// go task("Task 1")
	// go task("Task 2")
	// time.Sleep(time.Second * 2)
	// fmt.Println("Completed")

	// Main Example with Waitgroups
	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()

	fmt.Println("Completed")
}
