package main

import (
	"fmt"
	"math"
	"time"
)

func add(a, b int) int {
	return (a + b)
}

var i = 10

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	var p, q string
	var b time.Time
	b = time.Now()
	var a float64
	a = math.Pi
	fmt.Printf("the value of a is %f\n, and the value of b is %s\n", a, b)
	fmt.Printf("sum:%d\n", add(4, 5))
	fmt.Println(i)
	p, q = swap("love", "you")
	fmt.Printf("swapped strings are %s, %s\n", p, q)
}
