package main 

import "fmt"

func main() {

	fmt.Println("Check if a year is a leap year!")

	var year int
	fmt.Print("Enter a year: ")
	fmt.Scanf("%d", &year)

	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}
	fmt.Println("Check if a year is a leap year with another input!")
}