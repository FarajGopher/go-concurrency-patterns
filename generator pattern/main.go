package main

import (
	"fmt"
	"time"
)

func generator(max int) <-chan int {

	out := make(chan int)
	//we need to geenrate numbers to the max
	//but we dnnt want to consume infinite memory
	go func() {
		defer close(out)
		for i := 1; i <= max; i++ {
			out <- i
		}
	}()
	return out
}

func main() {
	max := 10
	start := time.Now()
	consume := generator(max)
	for v := range consume {
		fmt.Println("values: ",v)
	}
	fmt.Println("Goroutine+Channel Time:", time.Since(start))
}


