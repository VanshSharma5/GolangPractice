## Important URLs

1. https://pkg.go.dev
2. https://go.dev/doc/
3. https://go.dev/doc/effective_go {It's demonstrate how tp write better code in go}


# My Experiences about Go learning

1. The date format is Quite Weird because the weired way rather than the traditional %dd%mm:%yyyy style
```go
fmt.Println(presentTime)
// In date "01-02-2006" the 01 => Month, 02 => Day, 2006 => Year
fmt.Println(presentTime.Format("01-02-2006")) // its Faa... you have to put excatly 01-02-2006 else you are cooked
fmt.Println(presentTime.Format("01-02-2006 Monday")) // Faaa... again else you gonna cooked again
fmt.Println(presentTime.Format("01-02-2006 15:04:05 Msonday")) // Faaa... once again else you gonna cooked again
```
2. The memory allocation and deallocation is pretty nicer compare to C and C++. Automatic GC(Garbage Collection) is also impressive nowdays(2026).

3. Printing the Sliced part of list. Oh GOD its like Python but -ve index and jumps/step value both are not allowed
```go
fmt.Println("Only [1:3] slice on updated fruitList it ", fruitList[1:3])
fmt.Println("Only [1:] slice on updated fruitList it ", fruitList[1:])
fmt.Println("Only [:2] slice on updated fruitList it ", fruitList[:2])
```
