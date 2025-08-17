package main 

import "fmt"

func main() {
	fmt.Println("Generate prime numbers up to N")

	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	fmt.Printf("Prime numbers up to %d are:\n", n)
	for i := 2; i <= n; i++ {
		if IsPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}