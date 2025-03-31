package main

import "fmt"

func main() {

	/*

	   POINTERS--

	   So a pointer is a reference datatype which does'nt have its own value. it reads the value
	   of some other data type through a reference. you can also write the value of some datatype
	   if you have the reference of that data type. reference is basically an address in memory where
	   the value of the variable is stored. usually the memory address looks like 0x3432933.


	*/

	a, b := 12, "life"
	fmt.Println(a, b)

	q := &a //declares q and assigns it the reference(memory address) of a
	r := &b //declares r and assigns it the reference(memory address) of b

	fmt.Println(*q, *r) // if you just prinf q,r it will log memory addressess like 0x213342 0x432434
	// if you want to log the values stored in that adress, just put a * to access them.h
	*q = 16
	*r = "love"

	fmt.Println(a, b)

}
