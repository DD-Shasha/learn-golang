package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Power of a number calculator!")

	var base, exponent int
	fmt.Print("Enter the base number: ")
	fmt.Scan(&base)
	fmt.Print("Enter the exponent: ")
	fmt.Scan(&exponent)

	timeStart := time.Now()
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	timeLapsed := time.Since(timeStart)
	fmt.Printf("Time taken to calculate: %v\n", timeLapsed)
	fmt.Printf("%d raised to the power of %d is %d\n", base, exponent, result)
	
	fmt.Println("Now using recursion!")

	timeStartRec := time.Now()
	resultRec := power(base, exponent)
	timeLapsedRec := time.Since(timeStartRec)
	fmt.Printf("Time taken to calculate using recursion: %v\n", timeLapsedRec)
	fmt.Printf("%d raised to the power of %d using recursion is %d\n", base, exponent, resultRec)
}

func power(base, exponent int) int {
	if exponent == 0 {
		return 1
	}
	return base * power(base, exponent-1)
}