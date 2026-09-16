package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(`Welcome to 
Switch-Case`)

	source := rand.NewSource(time.Now().UnixNano())
	localRand := rand.New(source)

	diceNum := localRand.Intn(6) + 1

	switch diceNum {
	case 1:
		fmt.Println("Dice Rolls to 1")
	case 2:
		fmt.Println("Dice Rolls to 2")
		fallthrough
	case 3:
		fmt.Println("Dice Rolls to 3")
		fallthrough
	case 4:
		fmt.Println("Dice Rolls to 4")
		fallthrough
	case 5:
		fmt.Println("Dice Rolls to 5")
	case 6:
		fmt.Println("Dice Rolls to 6")
	default:
		fmt.Println("What is this Bro?")
	}

}
