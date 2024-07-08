package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	fibos := make(chan int)
	diffs := make(chan int)
	go Counter(100, naturals)
	go Squarer(naturals, squares)
	go Fibonacci(fibos)
	go Diff(squares, fibos, diffs)
	Printer(diffs)
}

func Counter(limit int, out chan<- int) {
	for x := 1; x < limit; x++ {
		out <- x
	}
	close(out)
}

func Squarer(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func Fibonacci(out chan<- int) {
	a, b := 1, 1
	for {
		out <- a
		a, b = b, a+b
	}
}

func Diff(a, b <-chan int, out chan<- int) {
	for {
		diff := 0
		select {
		case x := <-a:
			y := <-b
			diff = x - y
		case y := <-b:
			x := <-a
			diff = x - y
		}
		out <- diff
		if diff < 0 {
			close(out)
			return
		}
	}
}

func Printer(in <-chan int) {
	i := 1
	for x := range in {
		fmt.Println(i, x)
		i++
	}
}
