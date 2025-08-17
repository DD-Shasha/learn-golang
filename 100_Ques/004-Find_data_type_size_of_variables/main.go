package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Finding the size of all primitive types in Go:")

	var (
		b   bool
		i   int
		f   float64
		s   string
		c   complex128
	) 

	// Print the sizes of each variable type
	fmt.Printf("Size of bool: %d bytes\n", unsafe.Sizeof(b))
	fmt.Printf("Size of int: %d bytes\n", unsafe.Sizeof(i))
	fmt.Printf("Size of float64: %d bytes\n", unsafe.Sizeof(f))
	fmt.Printf("Size of string: %d bytes\n", unsafe.Sizeof(s))
	fmt.Printf("Size of complex128: %d bytes\n", unsafe.Sizeof(c))	
}