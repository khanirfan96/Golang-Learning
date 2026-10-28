package main

import (
	"fmt"
	"sync"
)

// Race Condition Code

// var count int

// func increment() {
// 	for i := 0; i < 10000; i++ {
// 		count++
// 	}
// }

// func main() {
// 	go increment()
// 	go increment()
// 	go increment()
// 	go increment()
// 	go increment()
// 	go increment()
// 	go increment()
// 	go increment()

// 	time.Sleep(time.Second)

// 	fmt.Println("Final Value of count is -----", count)
// }

// Solving Race condition with Mutex

var count int
var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		mu.Lock()
		count++
		mu.Unlock()
	}
}

func main() {

	var wg sync.WaitGroup

	wg.Add(8)

	go increment(&wg)
	go increment(&wg)
	go increment(&wg)
	go increment(&wg)
	go increment(&wg)
	go increment(&wg)
	go increment(&wg)
	go increment(&wg)

	wg.Wait()

	fmt.Println("Final Value of count is -----", count)
}
