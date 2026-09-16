package main

import "fmt"

func main() {
	fmt.Println("Hello there are Structs")
	// No Inheritence in Golang
	somebody := User{"mukesh ambani", "mukesh@jio.com", true, 54}
	fmt.Println(somebody)

	// its raw string enclose between `` back-tiks
	fmt.Println(`The %v gives onlu values while %+v gives the label with the values`)
	fmt.Printf("User details using %%+v are %+v\n", somebody)

	fmt.Printf("User Name is %v with email %v\n", somebody.Name, somebody.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    uint8
}
