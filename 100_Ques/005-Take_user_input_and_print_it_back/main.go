package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Take user input and print it back!")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter something: ")
	inp, _ := reader.ReadString('\n')
	inp = strings.TrimSpace(inp)
	fmt.Println("You entered:", inp)

	fmt.Println("Take user input and print it back with another input!")
	fmt.Print("Enter something: ")
	inp2, _ := reader.ReadString('\n')
	inp2 = strings.TrimSpace(inp2)
	fmt.Println("You entered:", inp2)

	fmt.Println("Take user input and print it back with two words!")
	fmt.Print("Enter two words: ")
	line, _ := reader.ReadString('\n')
	words := strings.Fields(line)
	if len(words) >= 2 {
		fmt.Println("You entered:", words[0], "and", words[1])
	} else {
		fmt.Println("Not enough words entered.")
	}

	fmt.Print("Enter a sentence: ")
	inp5, _ := reader.ReadString('\n')
	inp5 = strings.TrimSpace(inp5)
	fmt.Println("You entered:", inp5)

	fmt.Println("Take user input and print it back with two words (with spaces)!")
	fmt.Print("Enter two words with spaces: ")
	line2, _ := reader.ReadString('\n')
	words2 := strings.Fields(line2)
	if len(words2) >= 2 {
		fmt.Println("You entered:", words2[0], "and", words2[1])
	} else {
		fmt.Println("Not enough words entered.")
	}

	fmt.Println("Take user input and print it back with two words (with spaces and newlines)!")
	fmt.Print("Enter two words with spaces and newlines: ")
	line3, _ := reader.ReadString('\n')
	words3 := strings.Fields(line3)
	if len(words3) >= 2 {
		fmt.Println("You entered:", words3[0], "and", words3[1])
	} else {
		fmt.Println("Not enough words entered.")
	}
}
