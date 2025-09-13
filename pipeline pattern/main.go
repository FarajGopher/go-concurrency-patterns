package main

import "fmt"

func GenerateNumbers(num ...int) <-chan int {
	out := make(chan int)	
	go func(){
		for _,n := range num {
		out <- n
		}
		defer close(out)
	}()
	return out
}

func SquareNumbers(in <-chan int) <-chan int{
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		defer close(out)
	}()
	return out
}

func PrintNumbers(in <-chan int) {
	for n := range in {
		fmt.Println(n)
	}
}

func main() {
	numbers := GenerateNumbers(2, 3, 4)
	squares := SquareNumbers(numbers)
	PrintNumbers(squares)
}