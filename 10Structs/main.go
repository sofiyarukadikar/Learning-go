package main

import "fmt"

func main() {

	fmt.Println("Structs in go")

	//no inheritance in go
    sofiya := User{"Sofiya","sofiya@deshaw.com",true,20}
	fmt.Println(sofiya)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
