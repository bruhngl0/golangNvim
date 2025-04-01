package main

import (
	"fmt"
)

var x [3]string

func main() {

	x = [3]string{"love", "life", "live"}
	y := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(y)
	y[4] = 10
	fmt.Println(y)
	fmt.Println(x)

}
