package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")

	var ptr *int // the *int means pointer to int => a variable which is used to point an int variable
	fmt.Println("Value of ptr is ", ptr)
	fmt.Printf("The type of ptr is %T\n", ptr)

	num := 143

	// NOTE: If you learn pointes in C it not different at all
	var numptr = &num // the creation of pointer which point to variable num. The &num means the address of num
	fmt.Println("value of the pointer to num holds is ", numptr)
	fmt.Println("value of the pointer points to is ", *numptr) // *numptr just says i am the alias of the variable i am pointing to. Or in techinical words accesing the memory loation whose address the numptr holds

	// again the C suff. Upating the original num variable by the pointer numptr
	*numptr = *numptr * 2 // as the line 17, the "*numptr" is alias of "num" therefor it equivalent of "num = num * 2"
	fmt.Println("The Updated value of num is ", num, "The value of *numptr is ", *numptr)

}
