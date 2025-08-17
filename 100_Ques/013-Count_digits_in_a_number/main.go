package main 

import "fmt"

func main() {
	fmt.Println("Count digits in a number!")

	var count int
	var num int

	fmt.Print("Enter a number : ")
	fmt.Scan((&num))

	for i:=0; i<num; i++ {
		count++
		num /= 10
	}
	fmt.Println("Total digits:", count)
}