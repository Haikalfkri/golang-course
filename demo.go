package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Hello, World!")

	// ======== Data Types ===========

	// var x string = "Hello, World!"
	// var y uint32 = 10 // cannot store negative numbers
	// var z int32 = -10 // can store negative numbers
	// var a float32 = 3.14 // floating point number
	// var b float64 = 3.14 // floating point number
	// var c bool = true // boolean
	// var d bool = false // boolean
	// var e byte = 10 // byte is a single byte
	// var f rune = 'a' // rune is a single character

	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z)
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)
	// fmt.Println(f)

	// ========  Implicit and Explicit Type Declarations ===========

	// y := 10 // implicit type declaration and is automatically assigned the type of the value
	// y := 10.0 // float
	// y := false
	// y := uint8(10) // type cast
	// fmt.Println(y)
	// fmt.Printf("Type of y: %T\n", y)

	// Explicit Type Declaration
	// var x uint8
	// x = 5 // explicit type declaration
	// fmt.Println(x)

	// ====== Console Output =======

	// x := false
	// fmt.Println(x) // print the value of x
	// fmt.Printf("%T", x) // print formatted string with the type of x
	// y := fmt.Sprintln(x) // print the value of x and return the string
	// fmt.Println(y) // print the value of y

	// ====== Arithmetic Operations =======

	// fmt.Println(math.Min(4, 5))
	// fmt.Println(math.Max(4, 5))

	// x := "1234"
	// y, err := strconv.Atoi(x)
	// z, err := strconv.ParseInt(x, 10, 0)
	// fmt.Println(z, err)

	// ====== Conditions and Logical Operators =========

	// < > <= >= == != || && !

	// x := uint(8)
	// y := 10
	// z := int(x) == y
	// fmt.Println(z)

	// x := 2
	// if x < 3 {
	// 	fmt.Println("run")
	// } else if x > 4 {
	// 	fmt.Println("ok")
	// } else {
	// 	fmt.Println("error")
	// }

	// a := "h"
	// switch a {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// default:
	// 	fmt.Println("default")
	// }

	// switch a {
	// case "a", "b", "c":
	// 	fmt.Println("a b c")
	// default:
	// 	fmt.Println("default")
	// }

	// switch { // naked switch
	// case a == "a", a == "b":
	// 	fmt.Println("cek a and b")
	// default:
	// 	fmt.Println("default")
	// }

	// ============ Looping ============
	// for idx := 0; idx <= 10; idx++ {
	// 	fmt.Println(idx)
	// }

	// a := 0
	// for a < 10 { // while loop with for loop sintax
	// 	fmt.Println(a)
	// 	a++
	// }

	str := "hello world"
	// for idx := 0; idx < len(str); idx++ {
	// 	fmt.Printf("%c", str[idx])
	// }
	for _, char := range str {
		fmt.Printf("%c", char)
	}
}
