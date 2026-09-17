package main

import "fmt"

func main() {
	fmt.Println("Welcome from Loops")

	// It is slice not array
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	fmt.Println("Looping eash element via indexing")
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	fmt.Println("Looping with more better way i.e. using range keyword")
	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println("Printing via even more better way")
	for index, day := range days {
		fmt.Printf("Index is %v and %v is value\n", index, day)
	}

	roughValue := 1

	// For act like while when only single condition is passed
	for roughValue < 10 {

		if roughValue == 7 {
			break
		}

		if roughValue == 5 { // NOTE: This is while loop not for so we have to explicitely update the value of roughvalue else it became an infinite loop
			roughValue++
			continue
		}

		fmt.Println("Value is ", roughValue)
		roughValue++
	}

	faa := 1
customloop:
	fmt.Println("Now Faa no. ", faa)
	faa++
	if faa < 10 {
		goto customloop
	}

}
