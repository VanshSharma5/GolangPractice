package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("If Else in Go")

	reader := bufio.NewReader(os.Stdin)

	var result string

	fmt.Print("Please Enter a Number<Numeic Only>: ")
	input, _ := reader.ReadString('\n')

	loginCount, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 32)

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Something Else"
	} else {
		result = "Excatly 10 Login"
	}
	fmt.Println(result)

	if loginCount%2 == 0 {
		fmt.Println("Login Count is Even")
	} else {
		fmt.Println("Login Count is Odd")
	}

	// We can initilize the variable just right there inside the if statement then after a semcolon i can put the condition
	if num := 67; num < 10 {
		fmt.Println("Number is < 10")
	} else {
		fmt.Println("Number is >= 10")
	}
}
