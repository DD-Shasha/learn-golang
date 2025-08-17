package main 

import "fmt"

func main() {
	fmt.Println("Check if a number is even or odd!")
	
	var num checkEvenOrOdd
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)	

	fmt.Print(num.CheckIfEvenOrOdd())

}