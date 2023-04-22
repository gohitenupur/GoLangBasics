package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList =[]string{"Apple","Tomato","peach"}
	fmt.Println("Value of fruit list",fruitList)
	fmt.Printf("Type of fruitlist is %T \n", fruitList)

	fruitList =append(fruitList, "mango","banana")
	fmt.Println("Value of fruit list is", fruitList)

	//slices
	fruitList =append(fruitList[1:3])   // [1:] remove the first element ,[1:3] print 1 amd 2 
	fmt.Println(fruitList)

	highScorse :=make([]int ,4)
	highScorse[0] = 900	
	highScorse[1] = 200
	highScorse[2] = 300
    highScorse[3] = 400 // default alocation of the memory

	highScorse = append(highScorse, 444,456,567) // will append new values
	sort.Ints(highScorse) // provide sort function and many more functions
	fmt.Println(highScorse)
}