package main

import (
	"fmt"
	"math/big"

	// "math/rand" // comment it now because it conflict with the newxt import rand
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to My Maths")

	var intNum int = 2
	var floatNum float64 = 2.5

	fmt.Println("The sum is", intNum+int(floatNum)) // presion will gonna lost

	// // RANDOM NUMBER GENERATION With "math/rand"
	// source := rand.NewSource(time.Now().UnixNano())
	// localRand := rand.New(source)
	// // above 2lines are needed because rand.Seed() is depricated. So, we create a local rand object with local seed

	// fmt.Println(localRand.Intn(5))
	// fmt.Println(localRand.Intn(5))
	// fmt.Println(localRand.Intn(5))
	// fmt.Println(localRand.Intn(5))

	// RANDOM NUMBER GENERATION With "crypto/rand"
	myrandomnumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myrandomnumber)
}
