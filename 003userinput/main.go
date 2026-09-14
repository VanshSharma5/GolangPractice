package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the rating for our Pizza:")

	// comma ok || err ok

	input, _ := reader.ReadString('\n') // an anonymous placeholder to discard or ignore values that are syntactically required but not needed in your code
	fmt.Print("Thanks for rating", input)
	fmt.Printf("Type of input is %T\n", input)
}
