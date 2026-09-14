package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Print("Please rate out pizza between 1 to 5:")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Print("Thanks for rating: ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // strings is an package which contains some built-in public functions to perform soem operatiosn on strings

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating:", numRating+1)
	}
}
