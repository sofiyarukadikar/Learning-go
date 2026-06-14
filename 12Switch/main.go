package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switchcase in golang")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)
	switch diceNumber{
	case 1:
		fmt.Println("Dice value is 1 and you can now start the game")
	case 2:
		fmt.Println("Dice value is 2 and you can move 2")
	case 3:
		fmt.Println("Dice value is 3 and you can move 3")
	case 4:
		fmt.Println("Dice value is 4 and you can move 4")
	case 5:
		fmt.Println("Dice value is 5 and you can move 5")
	case 6:
		fmt.Println("Dice value is 6 and you can move 6")
	default:
		fmt.Println("wrong dice!!")
	}

	
}
