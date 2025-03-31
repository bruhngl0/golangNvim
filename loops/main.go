package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 0; i < 120; i++ {
		sum += i
	}

	fmt.Println(sum)

	/*we dont have while and do while loops in golang-- everything is done with in for loops.
	all the syntax for the while and dowhile can be implemented in for loops!! */

	diff := 1
	for diff < 1200 {
		diff += diff
	}

	fmt.Println(diff)

	/*

	  INFINITE LOOPS ---
		In case you are wondering why do we use infinite loops, we have some cases where we have to keep listening to
		the go routines, so when we recieve a signal, we break out of that loop--
		lets see one example for infinte loop first

		lets comment out this infinte loop first
		  inf :=1

		  for {
		    inf += inf
		  }

		this is an infinte loop where we havent defined any breakpoint for it. this is the pattern
		where we continuously check for the signals in go routines, now when we recieve a signal lets see
		how we get out of that infinte loop

	*/

	infC := 1

	for {
		if infC > 1201 {
			break
		}

		infC += infC
	}
	fmt.Println(infC)
}
