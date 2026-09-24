package main

import (
	"encoding/json"
	"fmt"
)

// data structure or say data model which is use to help in the seriliazation and decerialization of the struct into json and vice versa
// we can add extra info to the struct
type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	password string   `json:"-"`
	Tags     []string `json:"tags,omitempty` // don't add spaces before and after the comma(,) or it treate like a error|warning
}

func main() {
	fmt.Println("Welcome to JSON video")

	EncodeJson()
	DecodeJson()

}

func EncodeJson() {
	cour := []course{
		{"ReactJS SOme", 299, "LearnCodeOnline.in", "abc123!", []string{"web-dev", "js"}},
		{"Python", 199, "python.org", "py123!", []string{"python", "ai"}},
		{"go", 199, "go.org", "go123!", nil},
	}

	finalJson, err := json.MarshalIndent(cour, "", "  ") // Marshal => just plain json, while the MarshalIndent => add the indentation in the json while printing
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS SOme",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"Tags": [
			"web-dev",
			"js"
		]
	}
	`)

	var somecourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &somecourse)
		fmt.Printf("%#v\n", somecourse)
	} else {
		fmt.Println("JSON was not Valid")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and type is : %T\n", k, v, v)
	}

}
