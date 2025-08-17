package main 

import "fmt"

func main() {
	fmt.Println("Calculate GCD and LCM of two numbers")

	var num1, num2 int
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	gcd := GCD(num1, num2)
	lcm := LCM(num1, num2, gcd)

	fmt.Printf("GCD of %d and %d is: %d\n", num1, num2, gcd)
	fmt.Printf("LCM of %d and %d is: %d\n", num1, num2, lcm)
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b, gcd int) int {
	return (a * b) / gcd
}