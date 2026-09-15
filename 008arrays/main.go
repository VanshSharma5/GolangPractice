package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array in GO")

	var fruitList [4]string // the size is must along with data type

	fruitList[0] = "banana"
	// fruitList[1] = "orange"
	fruitList[2] = "Peach"
	fruitList[3] = "Guava"

	fmt.Println("Fruit list is ", fruitList) // OUTPUT: Fruit list is  [banana  Peach Guava]
	// NOTE: the 2 spaces between banana Peach contain a blank element in between
	fmt.Println("Length is fruit list is ", len(fruitList))

	var vegList = [5]string{"potato", "chili", "cucumber"}

	fmt.Println("Veggy list is ", vegList)
	fmt.Println("Veggy list size is ", len(vegList))

}
