package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)
	records := [][]string{
		{"1000000", "東京都", "千代田区", "-"},
		{"1020072", "東京都", "千代田区", "飯田橋"},
		{"1020082", "東京都", "千代田区", "一番町"}}

	writer.WriteAll(records)
	writer.Flush()
	file.Close()
}
