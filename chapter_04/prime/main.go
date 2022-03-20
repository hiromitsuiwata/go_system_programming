package main

import (
	"fmt"
	"math"
)

func primeNumber() chan int {
	result := make(chan int)
	go func() {
		result <- 2
		// 偶数は除外する
		for i := 3; i < 1000000; i += 2 {
			l := int(math.Sqrt(float64(i)))
			found := false
			for j := 3; j < l+1; j += 2 {
				// 2重ループで余りが0になる数が見つかったら素数ではない
				if i%j == 0 {
					found = true
					break
				}
			}
			// 余りが0になる数が見つからなかったら素数
			if !found {
				result <- i
			}
		}
		close(result)
	}()
	return result
}

func main() {
	pn := primeNumber()
	// 値が来るたびにforループが回る
	for n := range pn {
		fmt.Println(n)
	}
}
