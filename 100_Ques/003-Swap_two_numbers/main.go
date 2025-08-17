package main 

import "fmt"

func main(){
	fmt.Println("Swap two numbers with temp variable!")

	var a,b,c int
	a = 10
	b = 20
	fmt.Println("Before swapping: a =", a, ", b =", b)

	c = a
	a = b
	b = c
	fmt.Println("After swapping: a =", a, ", b =", b)
	fmt.Println("Swapping completed!")	

	fmt.Println("Swap two numbers without temp variable!")

	a = a + b
	b = a - b
	a = a - b
	fmt.Println("After swapping: a =", a, ", b =", b)
	fmt.Println("Swapping completed!")

	fmt.Println("Swap two numbers using bitwise XOR!")
	a = 10
	b = 20
	fmt.Println("Before swapping: a =", a, ", b =", b)
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println("After swapping: a =", a, ", b =", b)
	fmt.Println("Swapping completed!")

	fmt.Println("Swap two numbers using multiplication and division!")
	a = 10
	b = 20
	fmt.Println("Before swapping: a =", a, ", b =", b)
	a = a * b
	b = a / b
	a = a / b
	fmt.Println("After swapping: a =", a, ", b =", b)
	fmt.Println("Swapping completed!")

	fmt.Println("Swap two numbers after user input!")
	var x, y int
	fmt.Print("Enter first number: ")
	fmt.Scan(&x)
	fmt.Print("Enter second number: ")
	fmt.Scan(&y)
	fmt.Println("Before swapping: x =", x, ", y =", y)

	x = x + y
	y = x - y
	x = x - y
	fmt.Println("After swapping: x =", x, ", y =", y)
	fmt.Println("Swapping completed!")

	fmt.Println("Swap two numbers using pointers!")
	var p1, p2 *int
	p1 = &x
	p2 = &y
	*p1 = *p1 + *p2
	*p2 = *p1 - *p2
	*p1 = *p1 - *p2
	fmt.Println("After swapping using pointers: x =", x, ", y =", y)
	fmt.Println("Swapping completed!")

	fmt.Println("Swap two numbers using a function!")
	swap := func(a, b int) (int, int) {
		a = a + b
		b = a - b
		a = a - b
		return a, b
	}
	x, y = swap(x, y)
	fmt.Println("After swapping using function: x =", x, ", y =", y)
	fmt.Println("Swapping completed!")

	fmt.Println("Swap two numbers using a struct!")
	type Pair struct {
		A int
		B int		
	}
	pair := Pair{A: x, B: y}
	pair.A = pair.A + pair.B
	pair.B = pair.A - pair.B
	pair.A = pair.A - pair.B
	fmt.Println("After swapping using struct: x =", pair.A, ", y =", pair.B)
	fmt.Println("Swapping completed!")

	fmt.Println("Swap two numbers using a channel!")
	ch := make(chan int, 2)
	ch <- x
	ch <- y
	x = <-ch
	y = <-ch
	fmt.Println("After swapping using channel: x =", x, ", y =", y)
	fmt.Println("Swapping completed!")

	fmt.Println("Swap two numbers using a slice!")
	slice := []int{x, y}
	slice[0] = slice[0] + slice[1]
	slice[1] = slice[0] - slice[1]
	slice[0] = slice[0] - slice[1]
	x, y = slice[0], slice[1]
	fmt.Println("After swapping using slice: x =", x, ", y =", y)
	fmt.Println("Swapping completed!")

	fmt.Println("Swap two numbers using a map!")
	m := map[string]int{"x": x, "y": y}
	m["x"] = m["x"] + m["y"]
	m["y"] = m["x"] - m["y"]
	m["x"] = m["x"] - m["y"]
	x, y = m["x"], m["y"]
	fmt.Println("After swapping using map: x =", x, ", y =", y)
	fmt.Println("Swapping completed!")

	fmt.Println("Swap two numbers using a custom type!")
	type Number int
	var num1, num2 Number
	num1 = Number(x)
	num2 = Number(y)
	num1 = num1 + num2
	num2 = num1 - num2
	num1 = num1 - num2
	x, y = int(num1), int(num2)
	fmt.Println("After swapping using custom type: x =", x, ", y =", y)
	fmt.Println("Swapping completed!")

	fmt.Println("Swap two numbers using a closure!")
	swapClosure := func() (func() (int, int), func(int, int)) {
		var a, b int
		return func() (int, int) {
			return a, b
		}, func(x, y int) {
			a = x
			b = y
			a = a + b
			b = a - b
			a = a - b
		}
	}
	getValues, setValues := swapClosure()
	setValues(x, y)
	x, y = getValues()
	fmt.Println("After swapping using closure: x =", x, ", y =", y)
	fmt.Println("Swapping completed!")	

	fmt.Println("Swap two numbers using a recursive function!")
	var recursiveSwap func(int, int) (int, int)
	recursiveSwap = func(a, b int) (int, int) {
		if a == b {
			return a, b
		}
		a = a + b
		b = a - b
		a = a - b
		return a, b
	}
	x, y = recursiveSwap(x, y)
	fmt.Println("After swapping using recursive function: x =", x, ", y =", y)
	fmt.Println("Swapping completed!")

	fmt.Println("Swap two numbers using two input numbers !")
	var num11, num22 int
	fmt.Scan(&num11, &num22)
	fmt.Println("Before swapping: num1 =", num11, ", num2 =", num22)
	num11 = num11 + num22
	num22 = num11 - num22
	num11 = num11 - num22
	fmt.Println("After swapping: num1 =", num11, ", num2 =", num22)
	fmt.Println("Swapping completed!")

	fmt.Println("Swap two numbers using two input numbers !")
	var num111, num222 int
	fmt.Scan(&num111)
	fmt.Scan(&num222)
	fmt.Println("Before swapping: num1 =", num111, ", num2 =", num222)
	num111 = num111 + num222
	num222 = num111 - num222
	num111 = num111 - num222
	fmt.Println("After swapping: num1 =", num111, ", num2 =", num222)
	fmt.Println("Swapping completed!")

}