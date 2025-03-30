package main

import (
	"fmt"
	"math/cmplx"
)

var (
	toBe   bool       = false
	maxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {

	//--TYPE CONVERSIONS

	var i, j int = 3, 4
	var k float64 = float64(i*i + j*j)
	var l uint = uint(k)

	fmt.Println(i, j, k, l)

	fmt.Println(maxInt, z)
}
