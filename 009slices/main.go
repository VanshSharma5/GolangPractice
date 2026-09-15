package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slicing")

	var fruitList = []string{"Apple", "Banana", "Tomato"} // NOTE: it is slice not array, but somehow the ADT(abstract data type) is array
	fmt.Println("Data Before Appending")
	fmt.Println("data in fruitList ", fruitList)
	fmt.Printf("Type of fruitList %T\n", fruitList)
	fmt.Println("Length of fruitList", len(fruitList))

	fruitList = append(fruitList, "Mango", "Guava")
	fmt.Println("Data After Appending")
	fmt.Println("data in fruitList ", fruitList)
	fmt.Printf("Type of fruitList %T\n", fruitList)
	fmt.Println("Length of fruitList", len(fruitList))

	fmt.Println("Printing the Sliced part of list. Oh GOD its like Python but -ve index and jumps/step value both are not allowed")

	fmt.Println("Before updated fruitList using [1:] slice on it ", fruitList)
	fruitList = append(fruitList[1:]) // reallocate the new memory with appended data dynamically
	fmt.Println("New updated fruitList after [1:] slice on it ", fruitList)
	fmt.Println("Only [1:3] slice on updated fruitList it ", fruitList[1:3]) // NOTE: the last value 3 is not inclusive here means only the index 1, 2 are printed while the index 3 is excluded
	fmt.Println("Only [1:] slice on updated fruitList it ", fruitList[1:])
	fmt.Println("Only [:2] slice on updated fruitList it ", fruitList[:2])

	highScores := make([]int, 4) // make is used for dynamic memory allocation OR you say this way don't need "var"

	highScores[0] = 131
	highScores[1] = 756
	highScores[2] = 31
	highScores[3] = 231
	// highScores[4] = 231 // this statement could crash
	fmt.Println("highScores = ", highScores)

	highScores = append(highScores, 555, 777, 888) // again it dynamically reallocate new memory allocation for more data
	fmt.Println("highScores after appending = ", highScores)
	fmt.Println("check highScores is sorted or not = ", sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println("highScores after sorting = ", highScores)
	fmt.Println("check highScores is sorted or not = ", sort.IntsAreSorted(highScores))

	var subjects = []string{"Python", "Java", "C", "C++", "Golang", "JavaScript"}
	fmt.Println("subjects = ", subjects)
	index := 2
	subjects = append(subjects[:index], subjects[index+1:]...) // this line says take all before the "index" and take all next after "index" and put back into the subject
	fmt.Println("Subjects after manipulation", subjects)

}
