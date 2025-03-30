package main

import (
	"fmt"
	"math"
	"runtime"
)

func sqrt(x float64) string {

	if x < 0 {
		return sqrt(-x) + "i"
	} else {
		return fmt.Sprint(math.Sqrt(x))
	}

}

//now lets define a function that uses declares a variable inside the if/else condtion lexical scope
//you cant use this variable outside the scope
//if you try to use this declared z outside the scope it will give undefined error

func love(x, y int) string {
	if z := (x + y); z <= 20 {
		return fmt.Sprintf("i love you")
	}

	return fmt.Sprintf("fuck you")
}

//SWITCH CASES

func sw() {
	fmt.Println("my go runs on")

	switch os := runtime.GOOS; os {

	case "darwin":
		fmt.Println("OS X")

	case "linux":
		fmt.Println("linux")

	default:
		fmt.Printf("%s", os)
	}
}

func main() {
	fmt.Println(sqrt(15))
	fmt.Println(love(3, 4))
	sw()
}
