package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	ch := make(chan int, 10)

	arr := generateRandomArray(1000, 1000)

	for _, v := range arr {
		go func() {
			ch <- v
		}()
	}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("received value from channel - > %v", <-ch)
	}
}

//Returns integer array of length size between 0 and limit
func generateRandomArray(size, limit int) []int {
	if size <= 0 || limit < 0 {
		return []int{}
	}
	rand.Seed(time.Now().UnixNano())
	outp := make([]int, size)
	fmt.Println(len(outp))
	for i := 0; i < len(outp); i++ {
		outp[i] = rand.Intn(limit)
	}
	return outp
}
