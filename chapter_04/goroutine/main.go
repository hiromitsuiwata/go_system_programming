package main

import (
	"fmt"
	"time"
)

func sub() {
	fmt.Println("sub() is running")
	time.Sleep(time.Second)
	fmt.Println("sub() is done")
}

func main() {
	fmt.Println("start sub()")
	go sub()
	fmt.Println("main sleep")
	time.Sleep(2 * time.Second)

	tasks := make(chan string)
	go func() {
		tasks <- "hello"
	}()
	task := <-tasks
	fmt.Println(task)
}
