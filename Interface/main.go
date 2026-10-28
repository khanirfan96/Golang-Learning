package main

import (
	"fmt"
	"time"
)

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
	fmt.Println("Console Log----", message)
}

type FileLogger struct{}

func (f FileLogger) Log(message string) {
	fmt.Println("File Logger----", message)
}

type TimeLogger struct{}

func (t TimeLogger) Log(message string) {
	todayTime := time.Now().Format("2006-10-02 15:02:03")
	fmt.Println(todayTime, " ", message)
}

func ProcessTask(logger Logger) {
	logger.Log("Task Started")
	logger.Log("In Process")
	logger.Log("Task Completed")
}

func main() {
	consoleLogger := ConsoleLogger{}
	fileLogger := FileLogger{}
	timeLogger := TimeLogger{}

	fmt.Println("Console Logger----")
	ProcessTask(consoleLogger)

	fmt.Println("File Logger----")
	ProcessTask(fileLogger)

	fmt.Println("Time Logger----")
	ProcessTask(timeLogger)

}
