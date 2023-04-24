package main

import "fmt"

func main() {
	fmt.Println("Maps in goLan")
	 
	languages := make(map[string]string)
	languages["Go"] = "Go language"
	languages["Java"] = "Java language"	
	languages["Python"] = "Python language"
	languages["C"] = "C language"

	fmt.Println("Languages are :",languages)
	fmt.Println("Value is language[js]: ",languages["Go"])

	// loops in maps are 

	for key,value := range languages{
		fmt.Printf("Maps key is %v and value is %v \n ",key,value)
	}


}