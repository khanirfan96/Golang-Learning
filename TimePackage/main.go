package main

import (
	"fmt"
	"time"
)

func main() {

	time := time.Now()
	timeSecond := time.UnixNano()
	timeMillisecond := time.UnixMicro()
	fmt.Println("time---", time)
	fmt.Println("timeSecond---", timeSecond)
	fmt.Println("timeMillisecond---", timeMillisecond)

}
