package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Just run the fastapi server from the cli using "fastapi dev fastapi_server.py" or use uvicorn instead

func main() {
	fmt.Println("Welcome to Web Requests")

	performGetRequest()
	performPostJsonRequest()
	performFormRequest()
}

func performGetRequest() {
	const myurl = "http://localhost:8000/get"
	responce, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}
	defer responce.Body.Close()

	fmt.Println("Status Code: ", responce.StatusCode)
	fmt.Println("Content length", responce.ContentLength)

	bytedata, _ := io.ReadAll(responce.Body)
	var responceString strings.Builder // mutable string type
	byteCount, _ := responceString.Write(bytedata)

	fmt.Println("ByteCount: ", byteCount)
	fmt.Println(responceString.String())
}

func performPostJsonRequest() {
	const myurl = "http://127.0.0.1:8000/post"

	requestBody := strings.NewReader(`
		{
			"msg": "some",
			"name": "None",
			"price": 0,
			"availablity": "Unavailable"
		}
	`)

	responce, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer responce.Body.Close()

	bytedata, err := io.ReadAll(responce.Body)

	content := string(bytedata)

	fmt.Println(content)
}

func performFormRequest() {
	const myurl = "http://127.0.01:8000/post-form"

	data := url.Values{}

	data.Add("firstname", "patanahi")
	data.Add("lastname", "bhulgya")
	data.Add("email", "some@koi.com")

	responce, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}
	defer responce.Body.Close()

	databyte, _ := io.ReadAll(responce.Body)

	content := string(databyte)

	fmt.Println(content)
}
