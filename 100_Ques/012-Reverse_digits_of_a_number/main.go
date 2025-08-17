package main

import "fmt"

func main() {
	fmt.Println("Reverse digits of a number!")

	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	reversed := 0
	for num != 0 {
		remainder := num % 10
		reversed = reversed*10 + remainder
		num /= 10
	}

	fmt.Println("Reversed number:", reversed)
}