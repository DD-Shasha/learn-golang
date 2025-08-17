package main

import "fmt"

func main() {
	fmt.Println("all primitive types in Go:")
	var (
		b   bool
		i   int
		f   float64
		s   string
		c   complex128
	)

	fmt.Printf("bool: %v\n", b)
	fmt.Printf("int: %v\n", i)
	fmt.Printf("float64: %v\n", f)
	fmt.Printf("string: %v\n", s)
	fmt.Printf("complex128: %v\n", c)
}