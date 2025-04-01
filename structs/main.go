package main

import (
	"fmt"
)

func main() {

	type List struct {
		x int
		y int
	}
	v := List{10, 7}
	u := &v
	w := &List{10, 15}

	u.x = 1
	u.y = 2

	fmt.Println(List{15, 20})
	fmt.Println(v.x)
	fmt.Println(u)
	fmt.Println(*u)
	fmt.Println(u.x, u.y)
	fmt.Println(w)
	fmt.Println(*w)
	fmt.Println(w.x, w.y)

}
