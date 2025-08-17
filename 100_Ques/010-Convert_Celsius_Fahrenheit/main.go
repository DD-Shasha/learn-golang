package main 

import "fmt"

func main() {
	var celsius float64
	fmt.Print("Enter temperature in Celsius: ")
	fmt.Scan(&celsius)
	fahrenheit := (celsius * 9/5) + 32
	fmt.Printf("Temperature in Fahrenheit: %.2f\n", fahrenheit)
}