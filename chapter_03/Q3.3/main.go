package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("file.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	writer, err := zipWriter.Create("newfile.txt")
	if err != nil {
		panic(err)
	}

	reader := strings.NewReader("Hello World from strings.NewReader()")

	io.Copy(writer, reader)
}
