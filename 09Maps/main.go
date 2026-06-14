package main

import "fmt"

func main() {

	fmt.Println("maps in golang")

	//a map is used to store a key value pair -> map [key]value
	languages := make(map[string]string)
	//a make also works when you have not sepcified the size
	languages["JS"]="Javascript"
	languages["RB"]="Ruby"
	languages["GO"]="GOlang"
	fmt.Println("The map contains :",languages)

	//deleting a key

	delete(languages,"GO")
	fmt.Println("The map contains :",languages)

    //looping through the maps

	for key, value := range languages {
		fmt.Printf("For key %v value is  %v\n",key,value)
	}

}
