package main

import "fmt"

func main() {
	var fruitList = []string{}
	fruitList = append(fruitList, "Mango")
	fmt.Printf("fruit list is %T\n", fruitList)

	highScores := make([]int,4)
	highScores[0] = 234
	highScores[1] = 345
	highScores[2] = 456
	highScores[3] = 567
	
	fmt.Println(highScores)

	//part2
	//removing value from slice based on index
    var cources =[]string{"Math", "Science", "English","hindi"}
	fmt.Println(cources)
    cources = append(cources[:1],cources[2:]...)
	fmt.Println(cources)
	



}
