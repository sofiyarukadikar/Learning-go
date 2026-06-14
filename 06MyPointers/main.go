package main

import "fmt"

func main() {
	fmt.Println("Pointers")
//if you create a pointer and dont put any value in it its nil by deafult
var ptr *int
fmt.Println(ptr)

mynum := 23
//referencing is &
myptr := &mynum
fmt.Println(*myptr)
}
