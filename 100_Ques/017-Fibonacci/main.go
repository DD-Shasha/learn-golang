package main 

import "fmt"

func main() {
	fmt.Println("Fibonacci generator!")

	var n int
	fmt.Print("Enter the number of Fibonacci terms to generate: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Please enter a positive integer.")
		return
	}

	fibonacciSeries := GenerateFibonacci(n)
	fmt.Printf("Fibonacci series up to %d terms: %v\n", n, fibonacciSeries)
}

func GenerateFibonacci(n int) []int {
	fib := make([]int, n)
	fib[0] = 0
	if n > 1 {
		fib[1] = 1
		for i := 2; i < n; i++ {
			fib[i] = fib[i-1] + fib[i-2]
		}
	}
	return fib
}
