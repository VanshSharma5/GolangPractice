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

4. Golang do not primary supports OOPS. So, there is no such built-in inheritance(parent and child).

5. It is odd for me now that i can define struct after main and user it in main(Like use before its declaration/defination hits before its application)
```go
func main() {
	fmt.Println("Hello there are Structs")
	// No Inheritence in Golang
	somebody := User{"mukesh ambani", "mukesh@jio.com", true, 54}
	fmt.Println(somebody)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    uint8
}
```

6. The backtiks allows us to write multiline raw strings .
```go
func main() {
	fmt.Println(`Welcome to 
Switch-Case`)
}
```

7. The if-else have a inovative syntax to initilize and check the condition in the if statement 
```go
if num := 67; num < 10 { // the initilization of if is happen right before it using in if
	fmt.Println("Number is < 10")
} else {
	fmt.Println("Number is >= 10")
}
```
[!NOTE]: This `num` variable there is only avaliable inside the if and else blocks. So, it is not Exists outside the if-else-if block.

8. The `fallthrough` keyword is used to tell the switch case to fall through if we want while in other languages we use break to prevent fallthrough while here we said when we want the fallthrough to happen.
```go
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
```