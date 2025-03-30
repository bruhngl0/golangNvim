package main

import "fmt"

//DEFER-------

//so defer is a unique keyword in golang, whenever you use defer infront of any function,
//it pushes it to the stack (doesnt run it)
//it runs the rest of the program and after its finished, the main function goes to the stack and runs in a top popoff manner,
//here world is first pushed to the stack and then world2 is pushed.

// DEFER USECASE ----  defer is provided by go to prevent the memoryleaks.
//imaging you call a fn to run some reads and writes on a db. if a fn call has some error, the connection gives  error and
//doesnt reach the close function at the end of the code. so we define defer close fn at the top only, just in case there
//is some error, the defer function pops of the stack in the end to prevent the memory leak

//OUTPUT--WORKFLOW

//first defer--gets pushed in the stack
//second defer--gets pushed on top of the defer1

//runs normal hello

//main fn goes back to the stack and first calls the function on the top of the stack-- ie-- second defer
//then calls the defer1

//OUTPUT--
//hello
//world2
//world1

func main() {

	defer fmt.Println("world1")
	defer fmt.Println("world2")
	fmt.Println("hello")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("counting")
	fmt.Println("done")
}
