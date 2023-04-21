package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome :="welcome to user input"
	fmt.Println(welcome)

	// comma error /comma ok syntax
	// go to the bufio pkg ,os pkg  - pkg.go.dev/bufio

	reader :=bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name :")

	//comma ok //error ok syntax
	input,_ := reader.ReadString('\n')
	fmt.Println("Welcome ",input)
	fmt.Printf("Type of input is %T ",input)


}