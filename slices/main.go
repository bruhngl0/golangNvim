package main

import "fmt"

var a [5]int
var b []string

func main() {

	a = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	p := a[1:4]
	fmt.Println(p)

	var q = []int{2, 3, 4, 5, 9}
	fmt.Println(q)

	b = []string{"live", "love", "life"}
	fmt.Println(b)

	//slice of mixed datatype

	s := []struct {
		x int
		y bool
		z string
	}{
		{1, true, "aditya"},
		{2, false, "nanu"},
	}

	fmt.Println(s)
}
