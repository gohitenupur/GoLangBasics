package main

import "fmt"

func main() {
	fmt.Println("Stucts in golang ")
	fmt.Println("----------------------")

	user1 := User{"nupur","nupur@gmail.com","programmer",25}

	fmt.Println(user1)
	fmt.Printf("Details are :%+v\n", user1)

	fmt.Printf("individual Name and Email is :%v %v\n", user1.Name ,user1.Email)


}

type User struct {
	Name string 
	Email string
	Status string
	Age int

}