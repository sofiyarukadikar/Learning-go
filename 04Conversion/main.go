package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//type conversions
	fmt.Println("Welcome to pizza app")
	fmt.Println("Enter the rating for pizza from 1 to 5")
	reader := bufio.NewReader(os.Stdin)
	rating, _ := reader.ReadString('\n')
	fmt.Print("Thanks for rating us with a", rating)
	numrating , err := strconv.ParseFloat(strings.TrimSpace(rating),64)

	if err != nil{
		fmt.Println(err)
	} else
	{
		fmt.Printf("The type of rating is %T", numrating)
	}
}
