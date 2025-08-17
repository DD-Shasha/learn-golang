package main

import "fmt"

type MyClass string;

func (c MyClass) PrintExtended() {
	println("Extended: " + c)
}

type MyFavouriteTeacher struct {
	Name string
	Age  int
}	

func (t MyFavouriteTeacher) Print() {
	println("Name: " + t.Name + ", Age: " + fmt.Sprint(t.Age))	
}