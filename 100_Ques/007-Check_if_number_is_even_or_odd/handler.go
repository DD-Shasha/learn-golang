package main 

type checkEvenOrOdd int

func (n checkEvenOrOdd) CheckIfEvenOrOdd() string {
	if n%2 == 0 {
		return "The number is even."
	}
	return "The number is odd."
}