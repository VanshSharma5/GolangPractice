package main

import "fmt"

func main() {
	fmt.Println("Welcome to Functions")
	greeter()
	result := adder(3, 5)
	fmt.Println("Result of Student via adder is ", result)

	result2 := proAdder(3, 5, 1, 6, 8)
	fmt.Println("Result of Student via proAdder is ", result2)

	holdValueOne, holdValueRwo := ReturnsTwoValue()

	fmt.Println(holdValueOne, holdValueRwo)

}

// variadic functions takes variable number of parameters
func proAdder(values ...int) int {
	fmt.Printf("Type of values %T\n", values)
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}

// Fixed number of arguments
func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// Function returns multiple value use like => holdValueOne, holdValueRwo := ReturnsTwoValue()
func ReturnsTwoValue() (int, string) { // must specify the type of both returnee here with the parenthesis
	return 67, "Hellow Duniya"
}

func greeter() {
	fmt.Println("Welcome from Greater")
}
