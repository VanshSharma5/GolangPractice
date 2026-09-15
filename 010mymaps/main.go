package main

import "fmt"

func main() {
	fmt.Println("Welcome to MyMaps")

	languages := make(map[string]string) // [string] => type of key, type of value the key can associate => string

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List  of all languages ", languages)
	fmt.Println("JS is key of ", languages["JS"]) // accesing value by key

	// Deleting a key from map
	delete(languages, "RB") // removes key with associate value from the languages map
	fmt.Println("languages after deleting RB", languages)

	fmt.Println("Iterating over the map using loop")

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
