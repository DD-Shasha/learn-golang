package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	myVar := MyClass("Hello")
	myVar.PrintExtended()

	teacher := MyFavouriteTeacher{Name: "John Doe", Age: 30}
	teacher.Print()

	numbers := []int{1,2,3,4,5}

	for i,num := range numbers {
		if num%2 == 0 {
			fmt.Printf("Number %d at index %d is even\n", num, i)
		} else {
			fmt.Printf("Number %d at index %d is odd\n", num, i)
		}
	}

	

}
