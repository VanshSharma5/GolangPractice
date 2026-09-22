package main

import (
	"fmt"
	"io"
	"net/http"
)

const myurl string = "http://localhost/elsa.html"

func main() {

	fmt.Println("Welcome from URL in GoLang")
	responce, err := http.Get(myurl)

	checkNilError(err)
	defer responce.Body.Close() // caller's Responsibility to close the connection
	fmt.Printf("The Response data Type is %T\n", responce)

	databytes, err := io.ReadAll(responce.Body)

	checkNilError(err)
	content := string(databytes)

	fmt.Println(content)

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
