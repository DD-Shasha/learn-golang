package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Factorial Calculator!")

	var num int
	fmt.Print("Enter a number to calculate its factorial: ")
	fmt.Scan(&num)

	startTime := time.Now
	factorial := CalculateFactorial(num)
	elapsedTime := time.Since(startTime())
	fmt.Printf("Time taken to calculate factorial: %s\n", elapsedTime)
	fmt.Printf("The factorial of %d is %d\n", num, factorial)
}

func CalculateFactorial(n int) int {
	if(n == 0 || n == 1) {
		return 1
	}
	return n * CalculateFactorial(n-1)
}

