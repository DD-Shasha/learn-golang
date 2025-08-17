package main

import "fmt"

func main() {
	fmt.Println("Simple Calculator!	")
	var a, b float64
	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	fmt.Println("Choose an operation:")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	var choice int
	fmt.Print("Enter your choice (1-4): ")
	fmt.Scan(&choice)	

	switch choice {
	case 1:
		fmt.Printf("Result: %.2f + %.2f = %.2f\n", a, b, a+b)
	case 2:
		fmt.Printf("Result: %.2f - %.2f = %.2f\n", a, b, a-b)
	case 3:
		fmt.Printf("Result: %.2f * %.2f = %.2f\n", a, b, a*b)
	case 4:
		if b != 0 {
			fmt.Printf("Result: %.2f / %.2f = %.2f\n", a, b, a/b)
		} else {
			fmt.Println("Error: Division by zero is not allowed.")
		}
	default:
		fmt.Println("Invalid choice. Please select a valid operation (1-4).")
	}
}