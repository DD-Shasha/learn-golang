package main 

import "fmt"

func main() {
	fmt.Println("Check if a number is prime!")

	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)	

	if num <= 1 {
		fmt.Println("Please enter a number greater than 1.")
		return
	}

	if IsPrime(num) {
		fmt.Printf("%d is a prime number.\n", num)
	} else {
		fmt.Printf("%d is not a prime number.\n", num)
	}
}

func IsPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}