package main

import "fmt"

func main() {
	defer fmt.Println("")    // stack: "\n" <- top
	defer fmt.Print("!")     // stack "\n", "!" <- top
	defer fmt.Print("World") // stack "\n", "!", "World" <- top
	fmt.Print("Hello")       // This prints "Hello" to console
	fmt.Print(" ")           // THis prints " " to console
	// function reachs it end and there is no more statements are left to excute
	// SO, the "World" is poped out and printed
	// then "!" is printed
	// "\n" is printed
	// Hence, "Hello World!" is printed
	myDefr()
}

func myDefr() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("Number", i, "is Printed by defer statement")
		fmt.Println("Number", i, "is Printed by normal statement")
	}
}

// Faaaaa... you can understand the output the
// the defer of main are gonna executed when the myDefr() being completed, myDefr() is completed when its last defer being executed
