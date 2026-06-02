package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to Go programming!"
	//taking inputs
	reader := bufio.NewReader(os.Stdin)
	username, _ := reader.ReadString('\n')
	fmt.Printf("%s , %s", welcome, username)

	//typecasting
	num, _ := reader.ReadString('\n')
	fmt.Printf("The type of number is %T", num)

}
