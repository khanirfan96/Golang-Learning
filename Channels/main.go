package main

import (
	"fmt"
	"time"
)

// Unbuffered Channels

func worker(ch chan int) {
	fmt.Println("Worker Started")
	time.Sleep(2 * time.Second)
	ch <- 100
	fmt.Println("Worker value send")
}

func safeSend(ch chan int, value int) (err error) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Recovered from panic:", r)
			err = fmt.Errorf("could not send value: %v", r)
		}
	}()

	ch <- value
	return nil
}

func main() {
	// ch := make(chan int)
	// go worker(ch)
	// fmt.Println("Waiting for Value")
	// value := <-ch
	// fmt.Println("Value Recieved---", value)

	// Buffered Channel
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	close(ch)

	err := safeSend(ch, 4)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Channel ---", <-ch)
	fmt.Println("Channel ---", <-ch)
	fmt.Println("Channel ---", <-ch)

}
