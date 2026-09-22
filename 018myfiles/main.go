package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to Files")

	// some content to write in file
	content := `While Propositional Logic can represent simple facts, it lacks the expressiveness needed to represent relationships  or general rules about classes of objects [958084467, 51]. For example, in Propositional Logic, "Every human is mortal" would require a separate propositional variable for every individual human failing to capture the underlying universal relationship.`

	file, err := os.Create("./myfile.txt") // file is being created

	if err != nil { // if any error is encounterd i.e. err has something(whatever it is)
		panic(err) // still unknown to me
	}
	length, err := io.WriteString(file, content) // write some string content to the file

	if err != nil { // if any error is encounterd i.e. err has something(whatever it is)
		panic(err) // still unknown to me
	}
	fmt.Println("Length of the Writed content is ", length)
	defer file.Close() // Now were you write is is gonna be the last one

	readFile("./myfile.txt")
}

func readFile(filename string) {
	data, err := os.ReadFile(filename)

	checkNilError(err)
	fmt.Println("Text data from file \n======================[", filename, "]======================\n", data)
}

// It is recommended way to reduce the reduntant if blocks to check the error.
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
