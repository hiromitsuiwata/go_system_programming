package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])

	i, _ := strconv.Atoi(os.Args[1])

	timeChannel := time.After(time.Duration(i) * time.Second)
	<-timeChannel
	fmt.Printf("%d seconds elapsed", i)
}
