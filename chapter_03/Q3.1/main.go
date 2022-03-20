package main

import (
	"io"
	"os"
)

func main() {
	src, err := os.Open("test.csv")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create("test2.csv")
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	io.Copy(dst, src)
}
