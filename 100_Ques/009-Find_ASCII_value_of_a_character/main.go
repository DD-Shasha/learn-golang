package main 

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var ch string
	fmt.Print("Enter a character: ")
	fmt.Scan(&ch)
	fmt.Printf("The ASCII value of '%c' is: %d\n", ch, ch[0])


	var fullLine string
	fmt.Print("Enter a full line of text: ")
	reader := bufio.NewReader(os.Stdin)
	fullLine, _ = reader.ReadString('\n')
	trimmedLine := strings.TrimSpace(fullLine)
	fmt.Printf("The ASCII values of the characters in '%s' are:\n", trimmedLine)


	for i:=0; i<len(trimmedLine); i++ {
		for j:=0; j<i+1; j++ {
			fmt.Printf("%d", trimmedLine[i])
		}
		fmt.Println()
	}
}