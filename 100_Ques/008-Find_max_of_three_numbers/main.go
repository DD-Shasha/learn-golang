package main 

import "fmt"

func main() {
	a := 10
	b := 20
	c := 30
	fmt.Println("Finding the maximum of three numbers:")
	max := findMax(a, b, c)
	fmt.Printf("The maximum of %d, %d, and %d is: %d\n", a, b, c, max)
}

func findMax(x, y, z int) int {
	max := x
	if y > max {
		max = y
	}
	if z > max {
		max = z
	}
	return max
}