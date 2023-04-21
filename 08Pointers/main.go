package main

import "fmt"

func main() {
	fmt.Println("Pointer in GoLang")
	// pointer - passing memory address of any variable (actual value)
	// var ptr *int

	// fmt.Println("Value of Pointer is:",ptr)

	myNum :=23

	var ptr = &myNum

	fmt.Println("Value of Actual Pointer is :",ptr)  // memory address  0x00000000
	fmt.Println("Value of Pointer is:",*ptr)		// value contain value 23
	
	*ptr =*ptr+2
	fmt.Println("New value is:",myNum)        // value contain value 24



}