package main 

import "fmt"

func main() {

	fmt.Println("Sum of digits in a number !")

	var sum int
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)	

	for num != 0 {
		sum += num % 10
		num /= 10
	}
	fmt.Println("Total sum:", sum)
}