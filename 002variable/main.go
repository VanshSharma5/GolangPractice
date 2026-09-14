package main

import "fmt"

// NOTE: In GO making the first letter in the variable(identifire) name make it public.
// So, it can be used outside the current package
var PublicVariable string = "I an a Public variable because I am begin with Capital P"


// only way to declare global variables is using "var" keyword
var globalString = "Cooker" // i can ignore type but its recommend to put
var globalInt int = 8080 

// NOTE: in function scope it is not allow to unused variables
func main() {

	// only way to declare global variables is using "var" keyword
	var localString = "Cooker" // i can ignore type but its recommend to put
	fmt.Println(localString)
	fmt.Printf("localString is of type %T\n", localString)
	
	var localInt int = 8080 
	fmt.Println(localInt)
	fmt.Printf("localInt is of type %T\n", localInt)
	var localInt8 int8 = 80 
	fmt.Println(localInt8)
	fmt.Printf("localInt8 is of type %T\n", localInt8)
	var localInt64 int64 = 8080 
	fmt.Println(localInt64)
	fmt.Printf("localInt64 is of type %T\n", localInt64)
	var localUInt32 uint32 = 8080 
	fmt.Println(localUInt32)
	fmt.Printf("localUInt32 is of type %T\n", localUInt32)
	
	// Single variable
    message := "Hello, Go!" // Infers string
	fmt.Println(message)
	fmt.Printf("message is of type %T\n", message)
	// Multiple variables of the same or different types
    age, isDeveloper := 30, true // Infers int and bool
	fmt.Println(age, isDeveloper)
	fmt.Printf("age is of type %T and isDeveloper of type %T\n", age, isDeveloper)
	
	fmt.Println(globalInt)
	fmt.Printf("globalInt is of type %T\n", globalInt)
	fmt.Println(globalString)
	fmt.Printf("globalString is of type %T\n", globalString)

}