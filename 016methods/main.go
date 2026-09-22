package main

import "fmt"

type User struct {
	name   string
	email  string
	status bool
	age    uint8
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.status)
}

func (u User) NewMain() {
	u.email = "test@go.dev"

	fmt.Println("Email of this user is ", u.email)
}

func main() {
	fmt.Println("Hello there are Structs")

	somebody := User{"mukesh ambani", "mukesh@jio.com", true, 54}
	fmt.Println(somebody.status) // its accessible only in this scope as it is private
	somebody.GetStatus()
	somebody.NewMain()            // Even this line updates the email and displays the updated email first. But its not actualy updated
	fmt.Printf("%+v\n", somebody) // Because when we saw the actual object after the manipulation it ramain same i.e. unchange
	/* Why this is happening ?
	 * Because the object of user to the methods are passed by value so a copy of user is created and than updated.
	 * So, the updation inside the method do not modified the actual object.
	 *
	 * How to fix it?
	 * Pass the reference instead of value i.e. passed the address to the pointer so the changed are happend on the shared memory location and do persist
	 */

}
