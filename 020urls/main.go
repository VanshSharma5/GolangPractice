package main

import (
	"fmt"
	"net/url"
)

const myurl string = "http://localhost:80/elsa.html?age=22&gender=F"

func main() {

	fmt.Println("Welcome from URL in GoLang")
	fmt.Println(myurl)

	// parsing the URL
	result, err := url.Parse(myurl)
	checkNilError(err)

	fmt.Println(result.Scheme)   // gives the protocol i.e. http or https
	fmt.Println(result.Host)     // gives the hostname such as localhost or google.com
	fmt.Println(result.Path)     // gives the the path after the hostname such as /home/users
	fmt.Println(result.Port())   // gives the port number the request is make on if any. Note: It's a method
	fmt.Println(result.RawQuery) // gives the query parameters i.e. the string after the "?". NOTE: it's a string

	qparams := result.Query() // It returns the request parameters as a map object of map[string]string so they are pretty easy to use
	fmt.Printf("The type of query params are: %T\n", qparams)

	for key, val := range qparams {
		fmt.Println("the key is", key, "with the value", val)
	}

	// Becaue it's reference is gonna used and passed to the method calls so the reference of object is stored
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "localhost",
		Path:     "/auth",
		RawQuery: "user=admin",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("The newly generated Url is ", anotherUrl)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
