package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time")

	// getting current time
	presentTime := time.Now()

	fmt.Println(presentTime)
	// In date "01-02-2006" the 01 => Month, 02 => Day, 2006 => Year
	fmt.Println(presentTime.Format("01-02-2006"))                 // its Faa... you have to put excatly 01-02-2006 else you are cooked
	fmt.Println(presentTime.Format("01-02-2006 Monday"))          // Faaa... again else you gonna cooked again
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // Faaa... once again else you gonna cooked again

	// create time object from given data
	createDate := time.Date(2005, time.October, 3, 13, 32, 12, 0, time.Local) // the month and location must be give as like here
	fmt.Println(createDate)
	fmt.Println(createDate.Format("02-01-2006 Monday"))
}
