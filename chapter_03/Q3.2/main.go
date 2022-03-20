package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	rand := rand.Reader
	limitRand := io.LimitReader(rand, 1024)

	file, err := os.Create("rand.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(file, limitRand)
}
