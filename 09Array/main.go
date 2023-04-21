package main

import "fmt"

func main() {
	fmt.Println("Array is :")
	var NumList [5]int

	NumList[0]=1
	NumList[1]=2	
	NumList[2]=3
	NumList[3]=4
	NumList[4]=5

	fmt.Println("number list is",NumList)
	fmt.Println("length of list is",len(NumList))

	var vegList =[3]string{"potato","beans","mashrooms"};
	fmt.Println(" VegList is",vegList)
}